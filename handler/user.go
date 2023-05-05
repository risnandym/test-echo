package handler

import (
	"net/http"
	"strconv"
	"test-echo/model"
	"test-echo/service"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{userService}
}

// RegisterUser godoc
// @Summary Register a Full Access account (USER).
// @Description registering a user from full access.
// @Tags Auth
// @Param Body body model.CreateUserRequest true "the body to register a FULL ACCESS account (USER)"
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /users/ [post]
func (handler UserHandler) CreateUser(c echo.Context) error {
	req := model.CreateUserRequest{}
	if err := c.Bind(&req); err != nil {
		return err
	}
	user, err := handler.userService.CreateUser(req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Create user",
		"data":    user,
	})
}

// Get All User godoc
// @Summary Get All User.
// @Description get all user.
// @Tags User
// @Produce json
// @Success 200 {object} model.UserResponse
// @Router /users/ [get]
func (handler UserHandler) GetAllUser(c echo.Context) error {
	users, err := handler.userService.GetAllUser()
	//response handling with struct
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Get User Failed",
			Error:   err.Error(),
		})
	}
	//response handling with new map interface
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get All User",
		"data":    users,
	})
}

// GetUser godoc
// @Summary Get User.
// @Description get user.
// @Tags User
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path string true "user id"
// @Produce json
// @Success 200 {object} model.UserResponse
// @Router /users/{id} [get]
func (handler UserHandler) GetUserByID(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))

	user, err := handler.userService.GetUserById(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Get User by ID Failed",
			Error:   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Get User by ID",
		"data":    user,
	})
}

// UpdateUser godoc
// @Summary Update User.
// @Description update user.
// @Tags User
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Param id path string true "user id"
// @Param Body body model.UpdateUserRequest true "the body to register a FULL ACCESS account (USER)"
// @Produce json
// @Success 200 {object} model.UserResponse
// @Router /users/{id} [put]
func (handler UserHandler) UpdateUser(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))
	userRequest := model.UpdateUserRequest{}
	c.Bind(&userRequest)

	user, err := handler.userService.UpdateUser(userRequest, userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Update User Failed",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success Update User",
		"data":    user,
	})
}

// DeleteUser godoc
// @Summary Delete User
// @Description Delete User by id.
// @Tags User
// @Param Authorization header string true "Authorization. How to input in swagger : 'Bearer <insert_your_token_here>'"
// @Security BearerToken
// @Produce json
// @Param id path string true "user id"
// @Success 200 {object} string
// @Router /users/{id} [delete]
func (handler UserHandler) DeleteUser(c echo.Context) error {
	userId, _ := strconv.Atoi(c.Param("id"))
	err := handler.userService.DeleteUser(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, model.ErrorResponse{
			Code:    http.StatusBadRequest,
			Message: "Delete User Failed",
			Error:   err.Error(),
		})
	}
	return c.JSON(http.StatusOK, "Delete User Berhasil")
}
