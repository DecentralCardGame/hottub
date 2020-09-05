package types

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username   string     `json:"username" form:"username" query:"username"`
	Email      string     `gorm:"type:varchar(100);unique_index" json:"email" form:"email" query:"email"`
	Password   string     `json:"password" form:"password" query:"password"`
	Mnemonic   string     `json:"mnemonic" form:"mnemonic" query:"mnemonic"`
	CosmosUser CosmosUser `json:"cosmos_user" form:"cosmos_user" query:"cosmos_user" gorm:"ForeignKey:ID"`
	Token      string     `json:"token" form:"token" query:"token"`
	Admin      bool       `gorm:"default:false" json:"admin" query:"admin"`
}

func (u *User) HashPassword(plain string) (string, error) {
	if len(plain) == 0 {
		return "", errors.New("password should not be empty")
	}
	h, err := bcrypt.GenerateFromPassword([]byte(plain), bcrypt.DefaultCost)
	return string(h), err
}

func (u *User) ComparePassword(plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(plainText))
	return err == nil
}
