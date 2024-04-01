package main

import (
	"curso/devbook/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("API running!")

	router := router.Generate()
	log.Fatal(http.ListenAndServe(":5000", router))
}
