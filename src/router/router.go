package router

import (
	"curso/devbook/src/router/routes"

	"github.com/gorilla/mux"
)

func Generate() *mux.Router {
	r := mux.NewRouter()

	return routes.Config(r)
}
