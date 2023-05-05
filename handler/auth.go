package handler

import (
	"errors"
	"net/http"
	"test-echo/config"
	"test-echo/model"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
func getUserByEmail(e string) (*model.User, error) {
	var user model.User

	if err := config.DB.Model(model.User{}).Where("email = ?", e).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// LoginUser godoc
// @Summary Login.
// @Description Logging in to get jwt token to access admin or user api by roles.
// @Tags Auth
// @Param Body body LoginInput true "the body to login a user"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users/login [post]
func Login(c echo.Context) error {
	input := model.LoginUser{}
	_ = c.Bind(&input)
	email := input.Email
	pass := input.Password
	user, err := getUserByEmail(email)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Email",
			Error:   err.Error(),
		})
	}
	if user == nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "user not found",
			Error:   "user not found",
		})
	}
	if !CheckPasswordHash(pass, user.Password) {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Invalid Password",
		})
	}
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["name"] = user.Name
	claims["email"] = user.Email
	claims["age"] = user.Age
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	t, err := token.SignedString([]byte(config.Config("SECRET")))
	if err != nil {
		return echo.ErrUnauthorized
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Login Success",
		"data":    t,
	})
}
