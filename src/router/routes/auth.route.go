package routes

import (
	"api/src/controllers"
	"net/http"
)

var authRoutes = []Rota{
	{
		URI:    "/sign-in",
		Method: http.MethodPost,
		Action: controllers.Login,
		Auth:   false,
	},
}
