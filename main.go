package main

import (
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/swaggo/echo-swagger"
	"hottub/database"
	_ "hottub/docs"
	"hottub/handler"
)

var h *handler.Handler
var e *echo.Echo

// @title Crowdcontrol Hottub
// @version 1.0
// @description The API that controls authentication and user-management in CrowdControl

// @host hottub.crowdcontrol.network
// @BasePath /
func main() {
	db := database.New()
	initializeEcho()
	h = &handler.Handler{DB: db}
	registerRoutes()

	e.Logger.Fatal(e.Start(":1323"))
}

func registerRoutes() {
	e.GET("/", h.MainRoute)
	e.GET("/users", h.GetUsers)
	e.GET("/users/:id", h.GetUsersById)
	e.POST("/users", h.CreateUser)
	e.PUT("/users/:id", h.UpdateUser)
	e.DELETE("/users/:id", h.DeleteUser)
	e.PUT("/settings/:id", h.UpdateUserCosmosSettings)
	e.GET("/settings/:id", h.GetCosmosSettings)
	e.POST("/login", h.Login)
	e.POST("/register", h.Register)
	e.GET("/swagger/*", echoSwagger.WrapHandler)
}

func initializeEcho() {
	e = echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Use(middleware.Logger())
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(handler.Key),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Path() == "/login" || c.Path() == "/register" || c.Path() == "/" || c.Path() == "/swagger/*" {
				return true
			}
			return false
		},
	}))
	e.Use(middleware.Recover())
}
