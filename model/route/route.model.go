package routeModel

import (
	"net/http"
	errorModel "web-server/model/error"
)

type Path struct {
	MainPath string
	SubPath  string
}

type Route struct {
	Controller func(w http.ResponseWriter, r *http.Request) *errorModel.AppError
	Validation func(x interface{})
	Method     string
}
