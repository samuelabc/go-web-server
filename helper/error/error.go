package errorHelper

import (
	errorModel "web-server/model/error"
)

const (
	GENERAL_CODE = (iota + 1) * 1000
	HEALTH_CODE  = (iota + 1) * 1000
	ARTICLE_CODE = (iota + 1) * 1000
	ACCOUNT_CODE = (iota + 1) * 1000
)
const (
	VALIDATION_FAILED = iota + GENERAL_CODE
	DECODE_FAILED
	ENCODE_FAILED
	INVALID_REQUEST
	UNAUTHORIZED
	TOKEN_EXPIRED
	ACCESS_TOKEN_INVALID
	REFRESH_TOKEN_INVALID
)

func ErrValidation(err error) *errorModel.AppError {
	return &errorModel.AppError{
		Code:      VALIDATION_FAILED,
		Message:   "validation failed",
		ErrorData: err.Error(),
	}
}
func ErrDecode(err error) *errorModel.AppError {
	return &errorModel.AppError{
		Code:      DECODE_FAILED,
		Message:   "decode failed",
		ErrorData: err.Error(),
	}
}
func ErrEncode(err error) *errorModel.AppError {
	return &errorModel.AppError{
		Code:      ENCODE_FAILED,
		Message:   "encode failed",
		ErrorData: err.Error(),
	}
}
func ErrInvalidRequest(err error) *errorModel.AppError {
	return &errorModel.AppError{
		Code:      INVALID_REQUEST,
		Message:   "invalid request",
		ErrorData: err.Error(),
	}
}
func ErrUnauthorized(err error) *errorModel.AppError {
	return &errorModel.AppError{
		Code:      UNAUTHORIZED,
		Message:   "unauthorized",
		ErrorData: err.Error(),
	}
}
func ErrTokenExpired(err error) *errorModel.AppError {
	return &errorModel.AppError{
		Code:      TOKEN_EXPIRED,
		Message:   "token expired",
		ErrorData: err.Error(),
	}
}
func ErrAccessTokenInvalid(err error) *errorModel.AppError {
	return &errorModel.AppError{
		Code:      ACCESS_TOKEN_INVALID,
		Message:   "invalid access token",
		ErrorData: err.Error(),
	}
}
func ErrRefreshTokenInvalid(err error) *errorModel.AppError {
	return &errorModel.AppError{
		Code:      REFRESH_TOKEN_INVALID,
		Message:   "invalid refresh token",
		ErrorData: err.Error(),
	}
}

const (
	POST_HEALTH_FAILED = iota + HEALTH_CODE
)

func ErrPostHealth(err error) *errorModel.AppError {
	return &errorModel.AppError{
		Code:      POST_HEALTH_FAILED,
		Message:   "health post failed",
		ErrorData: err.Error(),
	}
}

const (
	FETCH_ARTICLE_FAILED = iota + ARTICLE_CODE
	LIST_ARTICLE_FAILED
	CREATE_ARTICLE_FAILED
	UPDATE_ARTICLE_FAILED
	DELETE_ARTICLE_FAILED
)

func ErrFetchArticle(err error) *errorModel.AppError {
	return &errorModel.AppError{
		Code:      FETCH_ARTICLE_FAILED,
		Message:   "fetch article failed",
		ErrorData: err.Error(),
	}
}
func ErrListArticle(err error) *errorModel.AppError {
	return &errorModel.AppError{
		Code:      LIST_ARTICLE_FAILED,
		Message:   "list article failed",
		ErrorData: err.Error(),
	}
}
func ErrCreateArticle(err error) *errorModel.AppError {
	return &errorModel.AppError{
		Code:      CREATE_ARTICLE_FAILED,
		Message:   "create article failed",
		ErrorData: err.Error(),
	}
}
func ErrUpdateArticle(err error) *errorModel.AppError {
	return &errorModel.AppError{
		Code:      UPDATE_ARTICLE_FAILED,
		Message:   "update article failed",
		ErrorData: err.Error(),
	}
}
func ErrDeleteArticle(err error) *errorModel.AppError {
	return &errorModel.AppError{
		Code:      DELETE_ARTICLE_FAILED,
		Message:   "delete article failed",
		ErrorData: err.Error(),
	}
}

const (
	FETCH_ACCOUNT_FAILED = iota + ACCOUNT_CODE
	LIST_ACCOUNT_FAILED
	CREATE_ACCOUNT_FAILED
	UPDATE_ACCOUNT_FAILED
	DELETE_ACCOUNT_FAILED
)

func ErrFetchAccount(err error) *errorModel.AppError {
	return &errorModel.AppError{
		Code:      FETCH_ACCOUNT_FAILED,
		Message:   "fetch account failed",
		ErrorData: err.Error(),
	}
}
func ErrListAccount(err error) *errorModel.AppError {
	return &errorModel.AppError{
		Code:      LIST_ACCOUNT_FAILED,
		Message:   "list account failed",
		ErrorData: err.Error(),
	}
}
func ErrCreateAccount(err error) *errorModel.AppError {
	return &errorModel.AppError{
		Code:      CREATE_ACCOUNT_FAILED,
		Message:   "create account failed",
		ErrorData: err.Error(),
	}
}
func ErrUpdateAccount(err error) *errorModel.AppError {
	return &errorModel.AppError{
		Code:      UPDATE_ACCOUNT_FAILED,
		Message:   "update account failed",
		ErrorData: err.Error(),
	}
}
func ErrDeleteAccount(err error) *errorModel.AppError {
	return &errorModel.AppError{
		Code:      DELETE_ACCOUNT_FAILED,
		Message:   "delete account failed",
		ErrorData: err.Error(),
	}
}
