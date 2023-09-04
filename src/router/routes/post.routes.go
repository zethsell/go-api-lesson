package routes

import (
	"api/src/controllers"
	"net/http"
)

var postRoutes = []Rota{
	{
		URI:    "/posts",
		Method: http.MethodPost,
		Action: controllers.PostStore,
		Auth:   true,
	},
	{
		URI:    "/posts",
		Method: http.MethodGet,
		Action: controllers.PostList,
		Auth:   true,
	},
	{
		URI:    "/posts/{postId}",
		Method: http.MethodGet,
		Action: controllers.PostShow,
		Auth:   true,
	},
	{
		URI:    "/posts/{postId}",
		Method: http.MethodPut,
		Action: controllers.PostUpdate,
		Auth:   true,
	},
	{
		URI:    "/posts/{postId}",
		Method: http.MethodDelete,
		Action: controllers.PostDelete,
		Auth:   true,
	},
	{
		URI:    "/users/{userId}/posts",
		Method: http.MethodGet,
		Action: controllers.ListUserPosts,
		Auth:   true,
	},
	{
		URI:    "/posts/{postId}/like",
		Method: http.MethodPost,
		Action: controllers.PostLike,
		Auth:   true,
	},
	{
		URI:    "/posts/{postId}/unlike",
		Method: http.MethodPost,
		Action: controllers.PostUnlike,
		Auth:   true,
	},
}
