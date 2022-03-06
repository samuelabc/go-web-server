package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func StartServer() error {
	var err error

	r := mux.NewRouter()

	err = Router(r)
	if err != nil {
		return fmt.Errorf("failed to prepare api route: %w", err)
	}
	if err := http.ListenAndServe("localhost:8080", r); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	log.Printf("Listening on %d\n", 8080)
	if err != nil {
		return fmt.Errorf("server is not serving: %w", err)
	}

	return nil
}
