package apiMain

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"

	errorModel "web-server/model/error"
	healthRoute "web-server/src/api/health"
	helloRoute "web-server/src/api/hello"
	contextHelper "web-server/src/helper/context"
	errorHelper "web-server/src/helper/error"
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
		case "/health":
			{
				var body healthRoute.PostHealthSchema
				payload = &body
			}
		case "/hello":
			{
				var body helloRoute.PostHelloSchema
				payload = &body
			}
		}
		err = json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			return errorHelper.DECODE_ERROR(err)
		}

		validate := validator.New()
		err = validate.Struct(payload)
		if err != nil {

			// this check is only needed when your code could produce
			// an invalid value for validation such as interface with nil
			// value most including myself do not usually have code like this.
			if _, ok := err.(*validator.InvalidValidationError); ok {
				return errorHelper.VALIDATION_ERROR(err)
			}

			for _, err := range err.(validator.ValidationErrors) {
				fmt.Println("validation error", err)
			}
			return errorHelper.VALIDATION_ERROR(err)
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

func PrepareRoute(r *mux.Router) error {
	routeMatrix := GetRouteMatrix()
	// routeInfo := routeMatrix[routeModel.Path{MainPath: r.URL.Path, SubPath: ""}]
	for k, v := range routeMatrix {
		r.Handle(k.MainPath, appHandler(v.Controller)).Methods(v.Method)
		fmt.Printf("path[%v] info[%v]\n", k, v)
	}

	r.Use(RequestBodyValidationMiddleware)
	return nil
}
