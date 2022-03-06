package routeModel

import (
	"net/http"
	errorModel "web-server/model/error"
)

var (
	GET    string = "GET"
	POST          = "POST"
	DELETE        = "DELETE"
)

type Path struct {
	MainPath string
	SubPath  string
	Method   string
}
type Route struct {
	Controller func(w http.ResponseWriter, r *http.Request) *errorModel.AppError
	// Method     string
	// Validation func(io.ReadCloser) error
}
