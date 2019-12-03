package hottub

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func mainRoute(c echo.Context) error {
	return c.String(http.StatusOK, "DecentralCardGame - Hottub")
}

func Init() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", mainRoute)
	e.GET("/users", GetUsers)
	e.GET("/users/:id", GetUsersById)
	e.POST("/users", CreateUser)
	e.PUT("/users/:id", UpdateUser)
	e.DELETE("/users/:id", DeleteUser)
	e.POST("/login", Login)

	e.Logger.Fatal(e.Start(":1323"))
}
