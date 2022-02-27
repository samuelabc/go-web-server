package main

import (
	"fmt"
	"os"
	server "web-server/src"
)

func main() {
	var err error

	err = server.StartServer()
	if err != nil {
		fmt.Println("%w", err)
		os.Exit(1)
	}
}
