package controllers

import (
	"net/http"
	"strconv"

	"submission/day2/dto"
	"submission/day2/lib/database"
	"submission/day2/models"

	"github.com/labstack/echo/v4"
)

func GetBookControllers(c echo.Context) error {
	response := new(dto.Response)
	books, e := database.GetBooks()

	if e != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = e.Error()
	} else {
		response.StatusCode = http.StatusOK
		response.Message = "Success Get Books"
		response.Data = books
	}

	return c.JSON(response.StatusCode, response)
}

func GetBookByIdControllers(c echo.Context) error {
	response := new(dto.Response)
	paramId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = err.Error()
		return c.JSON(response.StatusCode, response)
	}
	book, e := database.GetBookById(paramId)

	if e != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = e.Error()
	} else {
		response.StatusCode = http.StatusOK
		response.Message = "Success Get Book ID : " + c.Param("id")
		response.Data = book
	}

	return c.JSON(response.StatusCode, response)
}

func CreateBookController(c echo.Context) error {
	book := new(models.Books)
	response := new(dto.Response)
	c.Bind(book)

	bookCreated, e := database.CreateBook(book)

	if e != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Gagal Created Book"
	} else {
		response.StatusCode = http.StatusOK
		response.Message = "Success Created Book"
		response.Data = bookCreated
	}

	return c.JSON(response.StatusCode, response)
}

func UpdateBookByIdController(c echo.Context) error {
	book := new(models.Books)
	response := new(dto.Response)
	c.Bind(book)

	bookUpdated, e := database.UpdateBookById(book)

	if e != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = "Gagal Updated Book ID : " + c.Param("id")
	} else {
		response.StatusCode = http.StatusOK
		response.Message = "Success Updated Book ID : " + c.Param("id")
		response.Data = bookUpdated
	}

	return c.JSON(response.StatusCode, response)
}

func DeleteBookByIdController(c echo.Context) error {
	response := new(dto.Response)
	paramId, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = err.Error()
		return c.JSON(response.StatusCode, response)
	}
	_, e := database.DeleteBookById(paramId)

	if e != nil {
		response.StatusCode = http.StatusBadRequest
		response.Message = e.Error()
	} else {
		response.StatusCode = http.StatusOK
		response.Message = "Success Delete Book ID : " + c.Param("id")
	}

	return c.JSON(response.StatusCode, response)
}
