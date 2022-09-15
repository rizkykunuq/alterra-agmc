package routes

import (
	"submission/day2/controllers"

	"github.com/labstack/echo/v4"
)

func AllRoutes() *echo.Echo {
	e := echo.New()
	// route users
	e.GET("/users", controllers.GetUserControllers)
	e.GET("/users/:id", controllers.GetUserByIdControllers)
	e.POST("/users", controllers.CreateUserController)
	e.PUT("/users/:id", controllers.UpdateUserByIdController)
	e.DELETE("/users/:id", controllers.DeleteUserByIdController)
	// route books
	e.GET("/books", controllers.GetBookControllers)
	e.GET("/books/:id", controllers.GetBookByIdControllers)
	e.POST("/books", controllers.CreateBookController)
	e.PUT("/books/:id", controllers.UpdateBookByIdController)
	e.DELETE("/books/:id", controllers.DeleteBookByIdController)

	return e
}
