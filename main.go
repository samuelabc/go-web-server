package main

import (
	"fmt"
	"os"
	api "web-server/api"
)

func main() {
	var err error

	err = api.StartServer()
	if err != nil {
		fmt.Println("%w", err)
		os.Exit(1)
	}
}
