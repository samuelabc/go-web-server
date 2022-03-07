package accountController

import (
	"context"
	"encoding/json"
	"net/http"

	accountModel "web-server/model/account"

	errorHelper "web-server/helper/error"
	errorModel "web-server/model/error"

	"github.com/go-playground/validator/v10"
)

// ProfileStore defines database operations for a profile.
type AccountStore interface {
	Get(context.Context, *GetAccountRequest) (*accountModel.Account, error)
	// Register(*RegisterUserRequest) (*userModel.User, error)
	// Login(*RegisterUserRequest) (*userModel.User, error)
	// List(*ListUserRequest) (*[]userModel.User, error)
	// Update(*UpdateUserRequest) (*userModel.User, error)
	// Delete(*DeleteUserRequest) (*userModel.User, error)
}

// ArticleResource implements article management handler.
type AccountResource struct {
	Store AccountStore
}

// NewProfileResource creates and returns a profile resource.
func NewAccountResource(store AccountStore) *AccountResource {
	return &AccountResource{
		Store: store,
	}
}

type accountResponse struct {
	*accountModel.Account
}

func getAccountResponse(a *accountModel.Account) *accountResponse {
	return &accountResponse{
		Account: a,
	}
}

// func newArticleListResponse(p *[]models.Article) *[]articleResponse {
// 	// var res []articleResponse
// }

func (rs *AccountResource) get(w http.ResponseWriter, r *http.Request) *errorModel.AppError {
	var err error
	data := &GetAccountRequest{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		return errorHelper.ErrInvalidRequest(err)
	}

	validate := validator.New()
	err = validate.Struct(data)
	if err != nil {
		return errorHelper.ErrValidation(err)
	}

	account, err := rs.Store.Get(context.Background(), data)
	if err != nil {
		return errorHelper.ErrFetchArticle(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(getAccountResponse(account)); err != nil {
		return errorHelper.ErrEncode(err)
	}
	return nil
}

func (rs *AccountResource) register(w http.ResponseWriter, r *http.Request) *errorModel.AppError {
	// var err error

	// data := &RegisterUserRequest{}
	// if err := json.NewDecoder(r.Body).Decode(&data.Article); err != nil {
	// 	return errorHelper.ErrInvalidRequest(err)
	// }

	// validate := validator.New()
	// err = validate.Struct(data)
	// if err != nil {
	// 	return errorHelper.ErrValidation(err)
	// }

	// res, err := rs.Store.Create(data.Article)
	// if err != nil {
	// 	return errorHelper.ErrCreateArticle(err)
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// if err := json.NewEncoder(w).Encode(newArticleResponse(res)); err != nil {
	// 	return errorHelper.ErrEncode(err)
	// }
	return nil
}

func (rs *AccountResource) login(w http.ResponseWriter, r *http.Request) *errorModel.AppError {
	// var err error

	// data := &RegisterUserRequest{}
	// if err := json.NewDecoder(r.Body).Decode(&data.Article); err != nil {
	// 	return errorHelper.ErrInvalidRequest(err)
	// }

	// validate := validator.New()
	// err = validate.Struct(data)
	// if err != nil {
	// 	return errorHelper.ErrValidation(err)
	// }

	// res, err := rs.Store.Create(data.Article)
	// if err != nil {
	// 	return errorHelper.ErrCreateArticle(err)
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// if err := json.NewEncoder(w).Encode(newArticleResponse(res)); err != nil {
	// 	return errorHelper.ErrEncode(err)
	// }
	return nil
}

func (rs *AccountResource) list(w http.ResponseWriter, r *http.Request) *errorModel.AppError {
	// var err error
	// data := &ListArticleRequest{}
	// if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
	// 	return errorHelper.ErrInvalidRequest(err)
	// }

	// validate := validator.New()
	// err = validate.Struct(data)
	// if err != nil {
	// 	return errorHelper.ErrValidation(err)
	// }

	// articles, err := rs.Store.List(data)
	// if err != nil {
	// 	return errorHelper.ErrListArticle(err)
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// if err := json.NewEncoder(w).Encode(articles); err != nil {
	// 	return errorHelper.ErrEncode(err)
	// }
	return nil
}

func (rs *AccountResource) update(w http.ResponseWriter, r *http.Request) *errorModel.AppError {
	// var err error

	// data := &UpdateArticleRequest{}
	// if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
	// 	return errorHelper.ErrInvalidRequest(err)
	// }

	// validate := validator.New()
	// err = validate.Struct(data)
	// if err != nil {
	// 	return errorHelper.ErrValidation(err)
	// }

	// res, err := rs.Store.Update(data)
	// if err != nil {
	// 	return errorHelper.ErrUpdateArticle(err)
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// if err := json.NewEncoder(w).Encode(newArticleResponse(res)); err != nil {
	// 	return errorHelper.ErrEncode(err)
	// }
	return nil
}

func (rs *AccountResource) delete(w http.ResponseWriter, r *http.Request) *errorModel.AppError {
	// var err error

	// data := &DeleteArticleRequest{}
	// if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
	// 	return errorHelper.ErrInvalidRequest(err)
	// }

	// validate := validator.New()
	// err = validate.Struct(data)
	// if err != nil {
	// 	return errorHelper.ErrValidation(err)
	// }

	// res, err := rs.Store.Delete(data)
	// if err != nil {
	// 	return errorHelper.ErrDeleteArticle(err)
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	// if err := json.NewEncoder(w).Encode(newArticleResponse(res)); err != nil {
	// 	return errorHelper.ErrEncode(err)
	// }
	return nil
}
