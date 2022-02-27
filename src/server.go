package server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	apiMain "web-server/src/api"
)

func StartServer() error {
	var err error

	r := mux.NewRouter()

	err = apiMain.PrepareRoute(r)
	if err != nil {
		return fmt.Errorf("failed to prepare api route: %w", err)
	}

	err = http.ListenAndServe("localhost:8080", r)
	if err != nil {
		return fmt.Errorf("server is not serving: %w", err)
	}

	return nil
}
