package hottub

import "net/http"

type Error struct {
	Status int
	Title  string
	Detail string
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
