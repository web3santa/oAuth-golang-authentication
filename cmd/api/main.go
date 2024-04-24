package main

import (
	"fmt"
	"goauth/internal/server"
)

func main() {

	auth.newAuth()
	server := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("cannot start server: %s", err))
	}
}
