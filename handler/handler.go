package handler

import (
	"github.com/jinzhu/gorm"
	"hottub/types"
)

type (
	Handler struct {
		DB        *gorm.DB
		UserStore *types.UserStore
	}
)

func NewHandler(db *gorm.DB, us *types.UserStore) *Handler {
	return &Handler{
		DB:        db,
		UserStore: us,
	}
}

const (
	// Key (Should come from somewhere else).
	Key = "secret"
)
