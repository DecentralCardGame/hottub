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

type PublicUserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email"`
}

type PublicUsersRepsonse struct {
	Users []*PublicUserResponse `json:"users"`
}

func NewPublicUserResponse(u *User) *PublicUserResponse {
	r := new(PublicUserResponse)
	r.Email = u.Email
	u.Username = u.Username
	return r
}

func NewPublicUsersResponse(users []User) *PublicUsersRepsonse {
	r := new(PublicUsersRepsonse)
	r.Users = make([]*PublicUserResponse, 0)
	for _, u := range users {
		ur := new(PublicUserResponse)
		ur.Username = u.Username
		ur.Email = u.Email
		r.Users = append(r.Users, ur)
	}
	return r
}

type WelcomeResponse struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func NewWelcomeResponse() *WelcomeResponse {
	r := new(WelcomeResponse)
	r.Name = "DecentralCardGame - Hottub"
	r.Version = "v1"
}
