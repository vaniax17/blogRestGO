package post

import (
	"blogRest/src/core"
	"blogRest/src/db"
	"blogRest/src/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Create(c echo.Context) error {
	token := c.Request().Header.Get("Authorization")
	token = token[7:]
	claims, err := core.ValidateToken(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}
	username := core.GetUsername(claims)

	title := c.QueryParam("title")
	content := c.QueryParam("content")

	slug := Slug()

	post := &models.Post{
		Title:          title,
		Content:        content,
		AuthorUsername: username["username"],
		Slug:           slug,
	}

	db.DB.Create(post)
	return c.JSON(http.StatusOK, map[string]string{"message": "Post created successfully", "slug": slug})
}
