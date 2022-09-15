package controllers

import (
	"net/http"
	"strconv"

	"submission/day2/dto"
	"submission/day2/lib/database"
	"submission/day2/models"

	"github.com/labstack/echo/v4"
)

func GetUserControllers(c echo.Context) error {
	response := new(dto.Response)
	users, e := database.GetUsers()

	if e != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = e.Error()
	} else {
		response.StatusCode = http.StatusOK
		response.Message = "Success Get Users"
		response.Data = users
	}

	return c.JSON(response.StatusCode, response)
}

func GetUserByIdControllers(c echo.Context) error {
	response := new(dto.Response)
	paramId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = err.Error()
		return c.JSON(response.StatusCode, response)
	}
	user, e := database.GetUserById(paramId)

	if e != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = e.Error()
	} else {
		response.StatusCode = http.StatusOK
		response.Message = "Success Get User ID : " + c.Param("id")
		response.Data = user
	}

	return c.JSON(response.StatusCode, response)
}

func CreateUserController(c echo.Context) error {
	user := new(models.Users)
	response := new(dto.Response)
	c.Bind(user)

	userCreated, e := database.CreateUser(user)

	if e != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Gagal Created User"
	} else {
		response.StatusCode = http.StatusOK
		response.Message = "Success Created User"
		response.Data = userCreated
	}

	return c.JSON(response.StatusCode, response)
}

func UpdateUserByIdController(c echo.Context) error {
	user := new(models.Users)
	response := new(dto.Response)
	c.Bind(user)

	userUpdated, e := database.UpdateUserById(user)

	if e != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Gagal Updated User ID : " + c.Param("id")
	} else {
		response.StatusCode = http.StatusOK
		response.Message = "Success Updated User ID : " + c.Param("id")
		response.Data = userUpdated
	}

	return c.JSON(response.StatusCode, response)
}

func DeleteUserByIdController(c echo.Context) error {
	response := new(dto.Response)
	paramId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = err.Error()
		return c.JSON(response.StatusCode, response)
	}
	_, e := database.DeleteUserById(paramId)

	if e != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = e.Error()
	} else {
		response.StatusCode = http.StatusOK
		response.Message = "Success Delete User ID : " + c.Param("id")
	}

	return c.JSON(response.StatusCode, response)
}
