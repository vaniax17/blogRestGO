package user

import (
	"blogRest/src/core"
	"blogRest/src/db"
	"blogRest/src/models"
	"blogRest/src/validators"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Create(c echo.Context) error {

	username := c.QueryParam("username")
	valid := validators.IsUsername(username)
	if !valid {
		c.Logger().Error("Username is not valid")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Username is not valid min length 3 or max length is 30"})
	}
	password := c.QueryParam("password")
	email := c.QueryParam("email")
	valid = validators.IsEmail(email)
	if !valid {
		c.Logger().Error("Email is not valid")
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Email is not valid"})
	}
	hashedPassword := core.Hash(password)
	user := &models.User{
		Username: username,
		Password: hashedPassword,
		Email:    email,
	}

	if isCheckExists(user) {
		c.Logger().Error("User already exists")
		return c.JSON(http.StatusConflict, map[string]string{"error": "User already exists"})

	} else {
		db.DB.Create(user)
		c.Logger().Info("User created successfully")
		core.SetAndGenerateToken(username, c)
		return c.JSON(http.StatusOK, map[string]string{"message": "User created successfully"})
	}
}

func isCheckExists(user *models.User) bool {

	err := db.DB.Where("username = ? OR email = ?", user.Username, user.Email).First(user).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return false
	}
	return true

}

func Login(c echo.Context) error {

	username := c.QueryParam("username")
	password := c.QueryParam("password")
	user := &models.User{
		Username: username,
		Password: password,
	}

	exist := isCheckExists(user)
	if exist {
		core.GetWHereUser(user.Username, user)
		if core.Compare(user.Password, password) {
			token, err := core.GenerateToken(user.Username)
			if err != nil {
				return nil
			}
			c.Response().Header().Set("Authorization", "Bearer "+token)
			return c.JSON(http.StatusOK, map[string]string{"message": "Login successfully"})
		}
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}
	return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
}

func GetAll(c echo.Context) error {

	var usernames []string
	result := db.DB.Model(&models.User{}).Select("username").Find(&usernames)
	if result.Error != nil {
		err := result.Error
		return c.JSON(http.StatusInternalServerError, map[string]any{"error": err})
	}

	return c.JSON(http.StatusOK, map[string]any{"users": usernames})

}

func GetPosts(c echo.Context) error {

	username := c.QueryParam("username")
	user := &models.User{
		Username: username,
	}

	exists := isCheckExists(user)
	if !exists {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "User not found"})
	} else {

		core.GetWHereUser(username, user)
		if user.Posts == nil {
			return c.JSON(http.StatusOK, map[string]string{"user": user.Username, "posts": "No posts"})
		}

		return c.JSON(http.StatusOK, map[string]any{"user": user.Username, "posts": user.Posts})

	}

}

func ChangeUsername(c echo.Context) error {

	newUsername := c.QueryParam("new_username")

	token := c.Request().Header.Get("Authorization")
	if len(token) == 0 {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}
	token = token[7:]
	claims, err := core.ValidateToken(token)

	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": "Unauthorized"})
	}
	username := core.GetUsername(claims)

	user := &models.User{
		Username: newUsername,
	}
	exists := isCheckExists(user)
	valid := validators.IsUsername(newUsername)
	if !valid {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Username is not valid min length 3 or max length is 30"})
	}

	if !exists {
		db.DB.Model(models.User{}).Where("username = ?", username["username"]).Update("username", newUsername)
		core.SetAndGenerateToken(newUsername, c)
		return c.JSON(http.StatusOK, map[string]string{"message": "Username changed successfully"})
	}
	return c.JSON(http.StatusConflict, map[string]string{"error": "Username already exists"})
}
