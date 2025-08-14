package cores

import (
	"blogRest/src/db"
	"blogRest/src/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func SetAndGenerateToken(username string, c echo.Context) {

	token, err := GenerateToken(username)

	if err != nil {
		return
	}

	c.Response().Header().Set("Authorization", "Bearer "+token)

}

func GetWhereUser(username string, user *models.User) {
	db.DB.Preload("Posts").Where("username = ?", username).First(user)
}

func GetWherePost(slug string, post *models.Post) {
	db.DB.Where("slug = ?", slug).First(post)
}

func Slug() string {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return ""
	}
	return newUUID.String()
}
