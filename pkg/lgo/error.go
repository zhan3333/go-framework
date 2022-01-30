package lgo

import (
	"net/http"
)

type HTTPError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

func (H *HTTPError) Error() string {
	return H.Message
}

func NewHTTPError(code int, message ...string) error {
	he := &HTTPError{Code: code, Message: http.StatusText(code)}
	if len(message) > 0 {
		he.Message = message[0]
	}
	return he
}
