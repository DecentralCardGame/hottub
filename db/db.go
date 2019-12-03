package hottub

import (
	"github.com/jinzhu/gorm"
	hottub "hottub/types"
)

var db *gorm.DB
var err error

func Init() {
	db, err = gorm.Open("sqlite3", "test.db")
	if err != nil {
		panic("failed to connect database")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&hottub.User{})
}

func Manager() *gorm.DB {
	return db
}
