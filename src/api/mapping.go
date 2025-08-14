package api

import (
	"blogRest/src/post"
	"blogRest/src/user"

	"github.com/labstack/echo/v4"
)

func RoutesInit(e *echo.Echo) {
	e.GET("/user/login", user.Login)
	e.POST("/user/register", user.Create)
	e.GET("/user/all/usernames", user.GetAll)
	e.GET("/user/get/posts", user.GetPosts)
	e.PUT("/user/change/username", user.ChangeUsername)
	e.DELETE("/user/delete_my_account", user.DeleteMyAccount)
	e.POST("/posts", post.Create)
	e.GET("/posts/:slug", post.GetBySlug)
	e.GET("/posts", post.GetAll)
	e.DELETE("/posts/:slug", post.Delete)

}
