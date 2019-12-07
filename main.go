package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/gommon/log"
	_ "github.com/mattn/go-sqlite3"
	"hottub/handler"
	"hottub/types"
	"os"
)

var production bool

func main() {
	loadEnvironment()

	e := echo.New()
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

	var db *gorm.DB
	var err error
	if production {
		log.Printf("--- PRODUCTION MODE ACTIVE ---")
		db, err = gorm.Open("postgres", "host="+os.Getenv("PG_ADDRESS")+"port="+os.Getenv("PG_PORT")+" user="+os.Getenv("PG_USER")+" dbname="+os.Getenv("PG_DB")+" password="+os.Getenv("PG_PASSWORD"))
	} else {
		log.Printf("--- DEVELOPMENT MODE ACTIVE ---")
		db, err = gorm.Open("sqlite3", "main.db")
	}

	if err != nil {
		print(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&types.User{})
	db.AutoMigrate(&types.CosmosUser{})

	h := &handler.Handler{DB: db}

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

	e.Logger.Fatal(e.Start(":1323"))
}

func loadEnvironment() {
	if os.Getenv("ENVIRONMENT") == "PRODUCTION" {
		production = true
	} else {
		production = false
	}
}
