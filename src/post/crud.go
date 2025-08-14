package post

import (
	"blogRest/src/core"
	"blogRest/src/db"
	"blogRest/src/models"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
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

func checkExists(slug string) bool {
	post := &models.Post{}
	err := db.DB.Where("slug = ?", slug).First(post).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func GetBySlug(c echo.Context) error {
	slug := c.Param("slug")
	if !checkExists(slug) {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "Post not found"})
	}
	post := &models.Post{}
	db.DB.Where("slug = ?", slug).First(post)
	return c.JSON(http.StatusOK, map[string]any{"post": post})
}
