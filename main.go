package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	"hottub/handler"
	"hottub/routes"
	"hottub/types"
)

func main() {
	e := echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(handler.Key),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Path() == "/login" || c.Path() == "/register" {
				return true
			}
			return false
		},
	}))
	e.Use(middleware.Recover())

	db, err := gorm.Open("sqlite3", "main.db")
	if err != nil {
		print(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&types.User{})

	h := &handler.Handler{DB: db}

	e.GET("/", h.MainRoute)
	e.GET("/users", h.GetUsers)
	e.GET("/users/:id", h.GetUsersById)
	e.POST("/users", h.CreateUser)
	e.PUT("/users/:id", h.UpdateUser)
	e.DELETE("/users/:id", h.DeleteUser)
	e.POST("/login", h.Login)
	e.POST("/register", h.Register)

	e.Logger.Fatal(e.Start(":1323"))
}
