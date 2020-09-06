package types

import (
	"errors"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	gorm.Model
	Username string `json:"username" form:"username" query:"username"`
	Email    string `gorm:"type:varchar(100);unique_index" json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
	Mnemonic string `json:"mnemonic" form:"mnemonic" query:"mnemonic"`
	Token    string `json:"token" form:"token" query:"token"`
	Admin    bool   `gorm:"default:false" json:"admin" query:"admin"`
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

func (u *User) IsAdmin() bool {
	return u.Admin
}

// Store

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

// Get by ID
func (us *UserStore) GetByID(id uint) (*User, error) {
	var m User
	if err := us.db.First(&m, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

// Check whether user is admin
func (us *UserStore) GetByUsername(username string) (*User, error) {
	var m User

	if err := us.db.First(&m, &User{
		Username: username,
	}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}

	return &m, nil
}

// Check whether user is admin
func (us *UserStore) CheckUserAdmin(username string) (bool, error) {
	var m User

	if err := us.db.First(&m, User{
		Username: username,
	}).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return false, nil
		}
		return false, err
	}

	return m.IsAdmin(), nil
}

// Create user
func (us *UserStore) CreateNewUser(user *User) error {
	return us.db.Create(&user).Error
}
