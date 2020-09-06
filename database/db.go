package database

import (
	"fmt"
	"hottub/types"

	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func New() *gorm.DB {
	var db *gorm.DB
	var err error

	if os.Getenv("ENVIRONMENT") == "production" {
		fmt.Printf("--- PRODUCTION MODE ACTIVE ---\n")
		db, err = gorm.Open("postgres", "host="+os.Getenv("PG_ADDRESS")+" port="+os.Getenv("PG_PORT")+" user="+os.Getenv("PG_USER")+" dbname="+os.Getenv("PG_DB")+" password="+os.Getenv("PG_PASSWORD")+" sslmode=disable")
	} else {
		fmt.Printf("--- DEVELOPMENT MODE ACTIVE ---\n")
		db, err = gorm.Open("postgres", "host=localhost port=5432 user=postgres dbname=database sslmode=disable")
	}

	if err != nil {
		print(err.Error() + "\n")
		panic("failed to connect database\n")
	}

	db.DB().SetMaxIdleConns(3)
	db.LogMode(true)
	return db
}

func TestDB() *gorm.DB {
	db, err := gorm.Open("sqlite3", "./../realworld_test.database")
	if err != nil {
		fmt.Println("storage err: ", err)
	}
	db.DB().SetMaxIdleConns(3)
	db.LogMode(false)
	return db
}

func DropTestDB() error {
	if err := os.Remove("./../realworld_test.database"); err != nil {
		return err
	}
	return nil
}

func AutoMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&types.User{},
	)
}
