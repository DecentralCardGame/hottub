package hottub

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username   string     `json:"username" form:"username" query:"username"`
	Email      string     `gorm:"type:varchar(100);unique_index" json:"email" form:"email" query:"email"`
	Password   string     `json:"password" form:"password" query:"password"`
	CosmosUser CosmosUser `json:"cosmos_user" form:"cosmos_user" query:"cosmos_user"`
}
