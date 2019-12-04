package handler

import "github.com/jinzhu/gorm"

type (
	Handler struct {
		DB *gorm.DB
	}
)

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)
