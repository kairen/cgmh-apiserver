package httpwrapper

import (
	"net/http"
)

const (
	StatusSuccess             = http.StatusOK
	StatusBadRequest          = http.StatusBadRequest
	StatusUnauthorized        = http.StatusUnauthorized
	StatusForbidden           = http.StatusForbidden
	StatusNotFound            = http.StatusNotFound
	StatusInternalServerError = http.StatusInternalServerError
)

var msgs = map[int]string{
	StatusSuccess:             "Ok",
	StatusBadRequest:          "Bad Request",
	StatusUnauthorized:        "Unauthorized",
	StatusNotFound:            "Not Found",
	StatusForbidden:           "Forbidden",
	StatusInternalServerError: "Internal Server Error",
}

func getMsg(code int) string {
	return msgs[code]
}
