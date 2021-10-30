package main

import (
	"fmt"
	"go/src/config"
	"go/src/router"
	"log"
	"net/http"
)

func main() {

	config.Load()
	r := router.Generate()

	fmt.Printf("Listening on port : %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
