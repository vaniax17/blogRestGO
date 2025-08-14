package comment

import (
	cores "blogRest/src/core"
	"blogRest/src/db"
	"blogRest/src/models"
	"blogRest/src/post"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Create(c echo.Context) error {

	token := c.Request().Header.Get("Authorization")
	token = token[7:]
	claims, err := cores.ValidateToken(token)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}
	username := cores.GetUsername(claims)
	content := c.QueryParam("content")
	slugPost := c.Param("slug")
	slug := cores.Slug()
	comment := &models.Comment{
		Content:        content,
		AuthorUsername: username["username"],
		PostSlug:       slugPost,
		Slug:           slug,
	}
	db.DB.Create(comment)
	return c.JSON(http.StatusOK, map[string]any{"comment": comment})

}

func Get(c echo.Context) error {
	slugPost := c.Param("slug")
	exists := post.IsCheckExists(slugPost)
	if !exists {
		return c.JSON(http.StatusNotFound, map[string]string{"error": cores.PostNotFound})
	}
	comments := &[]models.Comment{}
	db.DB.Where("post_slug = ?", slugPost).Find(comments)
	if len(*comments) == 0 {
		return c.JSON(http.StatusOK, map[string]int{"comments": 0})
	}
	return c.JSON(http.StatusOK, map[string]any{"comments": comments})
}

func isExistsComment(slug string) bool {
	comment := &models.Comment{}
	err := db.DB.Where("slug = ?", slug).First(comment).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true
}

func GetConcrete(c echo.Context) error {
	slug := c.Param("slug")
	exists := isExistsComment(slug)
	if !exists {
		return c.JSON(http.StatusNotFound, map[string]string{"error": cores.CommentNotFound})
	}
	comment := &models.Comment{}
	db.DB.Where("slug = ?", slug).First(comment)
	return c.JSON(http.StatusOK, map[string]any{"comment": comment})
}
