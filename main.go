package main

import (
	"fmt"
	"hottub/handler"
	"hottub/types"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
)

var production bool
var h *handler.Handler
var e *echo.Echo
var db *gorm.DB
var err error

func main() {
	loadEnvironment()
	initializeEcho()
	initializeDB()
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
}

func initializeDB() {
	if production {
		fmt.Printf("--- PRODUCTION MODE ACTIVE ---")
		db, err = gorm.Open("postgres", "host="+os.Getenv("PG_ADDRESS")+" port="+os.Getenv("PG_PORT")+" user="+os.Getenv("PG_USER")+" dbname="+os.Getenv("PG_DB")+" password="+os.Getenv("PG_PASSWORD")+" sslmode=disable")
	} else {
		fmt.Printf("--- DEVELOPMENT MODE ACTIVE ---")
		db, err = gorm.Open("sqlite3", "main.db")
	}

	if err != nil {
		print(err.Error())
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&types.User{})
	db.AutoMigrate(&types.CosmosUser{})
}

func loadEnvironment() {
	if os.Getenv("ENVIRONMENT") == "production" {
		production = true
	} else {
		production = false
	}
}

func initializeEcho() {
	e = echo.New()
	e.Logger.SetLevel(log.ERROR)
	e.Use(middleware.Logger())
	e.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey: []byte(handler.Key),
		Skipper: func(c echo.Context) bool {
			// Skip authentication for and signup login requests
			if c.Path() == "/login" || c.Path() == "/register" || c.Path() == "/" {
				return true
			}
			return false
		},
	}))
	e.Use(middleware.Recover())
}
