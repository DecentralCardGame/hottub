package utils

import (
	"github.com/labstack/echo"
	"net/http"
)

type Error struct {
	Status int
	Title  string
	Detail string
}

type GeneralError struct {
	Errors map[string]interface{} `json:"errors"`
}

var ErrorParameterNotInteger = Error{
	Status: http.StatusBadRequest,
	Title:  "Not an Integer",
	Detail: "Parameter must be an integer!",
}

var ErrorCannotParseFields = Error{
	Status: http.StatusBadRequest,
	Title:  "Cannot parse fields",
	Detail: "The posted object cannot be parsed because fields don't match!",
}

func NewError(err error) GeneralError {
	e := GeneralError{}
	e.Errors = make(map[string]interface{})
	switch v := err.(type) {
	case *echo.HTTPError:
		e.Errors["body"] = v.Message
	default:
		e.Errors["body"] = v.Error()
	}
	return e
}

func AccessForbidden() GeneralError {
	e := GeneralError{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = "access forbidden"
	return e
}

func NotFound() GeneralError {
	e := GeneralError{}
	e.Errors = make(map[string]interface{})
	e.Errors["body"] = "resource not found"
	return e
}
