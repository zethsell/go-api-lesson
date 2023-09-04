package routes

import (
	. "api/src/middlewares"
	"github.com/gorilla/mux"
	"net/http"
)

type Rota struct {
	URI    string
	Method string
	Action func(w http.ResponseWriter, r *http.Request)
	Auth   bool
}

func Set(router *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, authRoutes...)
	routes = append(routes, postRoutes...)

	for _, route := range routes {
		if route.Auth {
			router.HandleFunc(route.URI, Auth(route.Action)).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI, route.Action).Methods(route.Method)
		}

	}

	return router
}
