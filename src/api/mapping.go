package api

import (
	"blogRest/src/user"

	"github.com/labstack/echo/v4"
)

func RoutesInit(e *echo.Echo) {
	e.GET("/user/login", user.Login)
	e.POST("/user/register", user.Create)
	e.GET("/user/all/usernames", user.GetAll)
	e.GET("/user/get/posts", user.GetPosts)

}
