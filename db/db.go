package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	hottub "hottub/types"
)

var db *gorm.DB
var err error

func Init() {
}
