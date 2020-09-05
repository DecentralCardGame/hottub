package types

import "github.com/labstack/echo/v4"

type UserLoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *UserLoginRequest) Bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

type UserRegisterRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Mnemonic string `json:"mnemonic" validate:"required"`
}

func (r *UserRegisterRequest) Bind(c echo.Context, u *User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.Username = r.Username
	u.Email = r.Email
	u.Mnemonic = r.Mnemonic
	h, err := u.HashPassword(r.Password)
	if err != nil {
		return err
	}
	u.Password = h
	return nil
}
