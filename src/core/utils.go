package core

import (
	"blogRest/src/db"
	"blogRest/src/models"

	"github.com/labstack/echo/v4"
)

func SetAndGenerateToken(username string, c echo.Context) {

	token, err := GenerateToken(username)

	if err != nil {
		return
	}

	c.Response().Header().Set("Authorization", "Bearer "+token)

}

func GetWHereUser(username string, user *models.User) {
	db.DB.Preload("Posts").Where("username = ?", username).First(user)
}
