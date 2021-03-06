package articleController

import (
	"encoding/json"
	"net/http"

	articleModel "web-server/model/article"

	errorHelper "web-server/helper/error"
	errorModel "web-server/model/error"

	"github.com/go-playground/validator/v10"
)

// ProfileStore defines database operations for a profile.
type ArticleStore interface {
	Get(*GetArticleRequest) (*articleModel.Article, error)
	List(*ListArticleRequest) (*[]articleModel.Article, error)
	Create(*CreateArticleRequest) (*articleModel.Article, error)
	Update(*UpdateArticleRequest) (*articleModel.Article, error)
	Delete(*DeleteArticleRequest) (*articleModel.Article, error)
}

// ArticleResource implements article management handler.
type ArticleResource struct {
	Store ArticleStore
}

// NewProfileResource creates and returns a profile resource.
func NewArticleResource(store ArticleStore) *ArticleResource {
	return &ArticleResource{
		Store: store,
	}
}

// func (rs *ArticleResource) router(r *mux.Router) {
// 	s = r.PathPrefix("/article").Subrouter()
// 	s.Handle("/", appHandler(executeRoute(v))).Methods(v.Method)
// 	r.Get("/", rs.fetch)
// 	r.Put("/", rs.update)
// 	return r
// }

// func (rs *ArticleResource) profileCtx(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		claims := jwt.ClaimsFromCtx(r.Context())
// 		p, err := rs.Store.Get(claims.ID)
// 		if err != nil {
// 			log(r).WithField("profileCtx", claims.Sub).Error(err)
// 			render.Render(w, r, ErrInternalServerError)
// 			return
// 		}
// 		ctx := context.WithValue(r.Context(), ctxProfile, p)
// 		next.ServeHTTP(w, r.WithContext(ctx))
// 	})
// }

type articleRequest struct {
	*articleModel.Article
}

func (d *articleRequest) Bind(r *http.Request) error {
	return nil
}

type articleResponse struct {
	*articleModel.Article
}

func newArticleResponse(p *articleModel.Article) *articleResponse {
	return &articleResponse{
		Article: p,
	}
}

// func newArticleListResponse(p *[]models.Article) *[]articleResponse {
// 	// var res []articleResponse
// }

func (rs *ArticleResource) get(w http.ResponseWriter, r *http.Request) *errorModel.AppError {
	var err error
	data := &GetArticleRequest{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return errorHelper.ErrInvalidRequest(err)
	}

	validate := validator.New()
	err = validate.Struct(data)
	if err != nil {
		return errorHelper.ErrValidation(err)
	}

	article, err := rs.Store.Get(data)
	if err != nil {
		return errorHelper.ErrFetchArticle(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(newArticleResponse(article)); err != nil {
		return errorHelper.ErrEncode(err)
	}
	return nil
}

func (rs *ArticleResource) list(w http.ResponseWriter, r *http.Request) *errorModel.AppError {
	var err error
	data := &ListArticleRequest{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return errorHelper.ErrInvalidRequest(err)
	}

	validate := validator.New()
	err = validate.Struct(data)
	if err != nil {
		return errorHelper.ErrValidation(err)
	}

	articles, err := rs.Store.List(data)
	if err != nil {
		return errorHelper.ErrListArticle(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(articles); err != nil {
		return errorHelper.ErrEncode(err)
	}
	return nil
}

func (rs *ArticleResource) create(w http.ResponseWriter, r *http.Request) *errorModel.AppError {
	var err error

	data := &CreateArticleRequest{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return errorHelper.ErrInvalidRequest(err)
	}

	validate := validator.New()
	err = validate.Struct(data)
	if err != nil {
		return errorHelper.ErrValidation(err)
	}

	res, err := rs.Store.Create(data)
	if err != nil {
		return errorHelper.ErrCreateArticle(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(newArticleResponse(res)); err != nil {
		return errorHelper.ErrEncode(err)
	}
	return nil
}

func (rs *ArticleResource) update(w http.ResponseWriter, r *http.Request) *errorModel.AppError {
	var err error

	data := &UpdateArticleRequest{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return errorHelper.ErrInvalidRequest(err)
	}

	validate := validator.New()
	err = validate.Struct(data)
	if err != nil {
		return errorHelper.ErrValidation(err)
	}

	res, err := rs.Store.Update(data)
	if err != nil {
		return errorHelper.ErrUpdateArticle(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(newArticleResponse(res)); err != nil {
		return errorHelper.ErrEncode(err)
	}
	return nil
}

func (rs *ArticleResource) delete(w http.ResponseWriter, r *http.Request) *errorModel.AppError {
	var err error

	data := &DeleteArticleRequest{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return errorHelper.ErrInvalidRequest(err)
	}

	validate := validator.New()
	err = validate.Struct(data)
	if err != nil {
		return errorHelper.ErrValidation(err)
	}

	res, err := rs.Store.Delete(data)
	if err != nil {
		return errorHelper.ErrDeleteArticle(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(newArticleResponse(res)); err != nil {
		return errorHelper.ErrEncode(err)
	}
	return nil
}
