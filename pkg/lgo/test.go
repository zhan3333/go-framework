package lgo

import (
	"encoding/json"
	"io"
	"strings"
)

func ToReader(body interface{}) io.Reader {
	b, _ := json.Marshal(body)
	return strings.NewReader(string(b))
}
