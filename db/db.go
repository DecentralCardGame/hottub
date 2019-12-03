package hottub

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	hottub "hottub/types"
)

var db *gorm.DB
var err error

func Init() {
	db, err = gorm.Open("sqlite3", "main.db")
	if err != nil {
		print(err.Error())
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&hottub.User{})
}

func Manager() *gorm.DB {
	return db
}
