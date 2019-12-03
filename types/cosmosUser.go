package hottub

import "github.com/jinzhu/gorm"

type CosmosUser struct {
	gorm.Model
	Mnemonic string
}
