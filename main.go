package main

import (
	"curso/devbook/src/config"
	"curso/devbook/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()

	router := router.Generate()

	fmt.Printf("API running on port %d!\n", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), router))
}
