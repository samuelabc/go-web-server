package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-pg/pg"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	articleController "web-server/api/article"
	helloRoute "web-server/api/hello"
	database "web-server/database"
	contextHelper "web-server/helper/context"
	errorHelper "web-server/helper/error"
	errorModel "web-server/model/error"
	routeModel "web-server/model/route"
)

func RequestBodyValidationMiddleware(next http.Handler) http.Handler {
	validationMiddleware := func(w http.ResponseWriter, r *http.Request) *errorModel.AppError {
		var err error
		var payload interface{}

		fmt.Println("path", r.URL.Path)

		// routeMatrix := GetRouteMatrix()
		// routeInfo := routeMatrix[routeModel.Path{MainPath: r.URL.Path, SubPath: ""}]
		// schema := reflect.TypeOf(routeInfo.Validation)
		// payload := routeInfo.Validation.(schema)
		// fmt.Println("payload type", reflect.TypeOf(payload))
		// payload := healthRoute.PostHealthSchema{}

		switch r.URL.Path {
		case "/hello":
			{
				var body helloRoute.PostHelloSchema
				payload = &body
			}
		}
		err = json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			return errorHelper.ErrDecode(err)
		}

		validate := validator.New()
		err = validate.Struct(payload)
		if err != nil {

			// this check is only needed when your code could produce
			// an invalid value for validation such as interface with nil
			// value most including myself do not usually have code like this.
			if _, ok := err.(*validator.InvalidValidationError); ok {
				return errorHelper.ErrValidation(err)
			}

			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println("validation error", err)
			}
			return errorHelper.ErrValidation(err)
		}

		fmt.Println("jsonbody", payload)
		ctx := context.WithValue(r.Context(), contextHelper.ContextKeyJSONPayload, payload)
		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r.WithContext(ctx))
		return nil
	}
	return appHandler(validationMiddleware)
}

type appHandler func(http.ResponseWriter, *http.Request) *errorModel.AppError

func (fn appHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)

		errResponse := errorModel.ErrorResponse{
			Success: false,
			Error:   *e,
		}
		fmt.Println("errResponse", errResponse)

		str, err := json.Marshal(&errResponse)
		if err != nil {
			fmt.Println("encode error", err)
		}
		w.Write(str)
	}
}

func executeRoute(routeInfo routeModel.Route) func(w http.ResponseWriter, r *http.Request) *errorModel.AppError {
	exucuteRouteFunc := func(w http.ResponseWriter, r *http.Request) *errorModel.AppError {
		if appErr := routeInfo.Controller(w, r); appErr != nil {
			return appErr
		}
		return nil
	}
	return exucuteRouteFunc
}

type API struct {
	Article *articleController.ArticleResource
}

// NewAPI configures and returns application API.
func NewAPI(db *pg.DB) (*API, error) {
	articleStore := database.NewArticleStore(db)
	article := articleController.NewArticleResource(articleStore)

	api := &API{
		Article: article,
	}
	return api, nil
}

func Router(r *mux.Router) error {
	// logger := logging.NewLogger()

	db, err := database.DBConn()
	if err != nil {
		// logger.WithField("module", "database").Error(err)
		return err
	}

	appAPI, err := NewAPI(db)
	if err != nil {
		// logger.WithField("module", "app").Error(err)
		return err
	}

	routeMatrix := GetRouteMatrix(appAPI)
	// routeInfo := routeMatrix[routeModel.Path{MainPath: r.URL.Path, SubPath: ""}]
	for k, v := range routeMatrix {
		fullPath := fmt.Sprint(k.MainPath, k.SubPath)
		r.Handle(fullPath, appHandler(executeRoute(v))).Methods(k.Method)
		fmt.Println(k.Method, fullPath)
	}

	// r.Use(logging.NewStructuredLogger(logger))
	// r.Use(RequestBodyValidationMiddleware)
	return nil
}
