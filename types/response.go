package types

import (
	"hottub/utils"
)

type UserLoginResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Mnemonic string `json:"mnemonic" form:"mnemonic" query:"mnemonic"`
	Token    string `json:"token" form:"token" query:"token"`
}

func NewUserLoginResponse(u *User) *UserLoginResponse {
	r := new(UserLoginResponse)
	r.Username = u.Username
	r.Email = u.Email
	r.Mnemonic = u.Mnemonic
	r.Token = utils.GenerateJWT(u.ID)
	return r
}
