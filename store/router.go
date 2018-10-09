package store

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

var controller = &Controller{Repository: Repository{}}

var routes = []Route{
	Route{
		Name:        "Login",
		Method:      "POST",
		Pattern:     "/api/login",
		HandlerFunc: controller.login,
	},
	Route{
		Name:        "GetAllMovie",
		Method:      "GET",
		Pattern:     "/api/getAllMovie",
		HandlerFunc: controller.getAllMovie,
	},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.HandlerFunc
		log.Println(route.Name)
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
