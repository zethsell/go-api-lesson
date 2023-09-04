package routes

import (
	"api/src/controllers"
	"net/http"
)

var userRoutes = []Rota{
	{
		URI:    "/users",
		Method: http.MethodGet,
		Action: controllers.UserList,
		Auth:   true,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodGet,
		Action: controllers.UserShow,
		Auth:   false,
	},
	{
		URI:    "/users",
		Method: http.MethodPost,
		Action: controllers.UserStore,
		Auth:   false,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodPut,
		Action: controllers.UserUpdate,
		Auth:   true,
	},
	{
		URI:    "/users/{userId}",
		Method: http.MethodDelete,
		Action: controllers.UserDelete,
		Auth:   false,
	},
	{
		URI:    "/users/{userId}/follow",
		Method: http.MethodPost,
		Action: controllers.Follow,
		Auth:   true,
	},
	{
		URI:    "/users/{userId}/unfollow",
		Method: http.MethodPost,
		Action: controllers.Unfollow,
		Auth:   true,
	},
	{
		URI:    "/users/{userId}/followers",
		Method: http.MethodGet,
		Action: controllers.Followers,
		Auth:   true,
	},
	{
		URI:    "/users/{userId}/following",
		Method: http.MethodGet,
		Action: controllers.Following,
		Auth:   true,
	},
	{
		URI:    "/users/{userId}/password",
		Method: http.MethodPost,
		Action: controllers.PasswordChange,
		Auth:   true,
	},
}
