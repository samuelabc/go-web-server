package errorHelper

import (
	errorModel "web-server/model/error"
)

var ErrorCodes map[string]errorModel.AppError

func VALIDATION_ERROR(err error) *errorModel.AppError {
	return &errorModel.AppError{
		Code:      1001,
		Message:   "validation failed",
		ErrorData: err.Error(),
	}
}
func DECODE_ERROR(err error) *errorModel.AppError {
	return &errorModel.AppError{
		Code:      1002,
		Message:   "decode failed",
		ErrorData: err.Error(),
	}
}
func ENCODE_ERROR(err error) *errorModel.AppError {
	return &errorModel.AppError{
		Code:      1003,
		Message:   "encode failed",
		ErrorData: err.Error(),
	}
}
func POST_HEALTH_ERROR(err error) *errorModel.AppError {
	return &errorModel.AppError{
		Code:      2001,
		Message:   "health post failed",
		ErrorData: err.Error(),
	}
}
