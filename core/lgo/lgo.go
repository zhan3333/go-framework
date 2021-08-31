package lgo

import (
	"fmt"
	"net/http"
)

type HTTPError struct {
	Code     int         `json:"-"`
	Message  interface{} `json:"message"`
	Internal error       `json:"-"`
}

func (H *HTTPError) Error() string {
	if H.Internal == nil {
		return fmt.Sprintf("code=%d, message=%s", H.Code, H.Message)
	}
	return fmt.Sprintf("code=%d, message=%s, internal=%s", H.Code, H.Message, H.Internal)
}

func (H *HTTPError) Unwrap() error {
	return H.Internal
}

func NewHTTPError(code int, message ...interface{}) *HTTPError {
	he := &HTTPError{Code: code, Message: http.StatusText(code)}
	if len(message) > 0 {
		he.Message = message[0]
	}
	return he
}
