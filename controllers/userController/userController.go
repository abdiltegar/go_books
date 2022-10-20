package userController

import (
	"github.com/labstack/echo/v4"
	"learn_orm/dto"
	"learn_orm/services/user"
	"net/http"
	"strconv"
)

type userController struct {
	userService user.UserService
}

func NewUserController(userService user.UserService) *userController {
	return &userController{
		userService,
	}
}

// get all users
func (ctrl *userController) GetUsersController(c echo.Context) error {
	users, err := ctrl.userService.GetAll()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func (ctrl *userController) GetUserController(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 16, 64)

	user, err := ctrl.userService.GetById(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get user",
		"user":    user,
	})
}

// create user
func (ctrl *userController) CreateUserController(c echo.Context) error {
	user := dto.DTOUserReq{}
	c.Bind(&user)

	err := ctrl.userService.Create(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
	})
}

// delete user by id
func (ctrl *userController) DeleteUserController(c echo.Context) error {
	id, _ := strconv.ParseInt(c.Param("id"), 16, 64)

	err := ctrl.userService.Delete(uint(id))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete user",
	})
}

// update user by id
func (ctrl *userController) UpdateUserController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 16, 32)

	user := dto.DTOUserReq{}
	c.Bind(&user)

	err := ctrl.userService.Update(uint(id), user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
	})
}

func (ctrl *userController) LoginController(c echo.Context) error {

	user := dto.DTOUserReq{}
	c.Bind(&user)

	res, err := ctrl.userService.Login(user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "login success",
		"data":    res,
	})
}
