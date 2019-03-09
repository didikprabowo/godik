package main

import (
	"fmt"
	"gitlab.com/didik/godik/system"
	"net/http"
)

func main() {
	router := config.RegisterRoutes()
	fmt.Println(http.ListenAndServe(":8080", router))
}
