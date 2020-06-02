package tool

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/skip2/go-qrcode"
	"log"
)

func Dump(data interface{}) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	jsonEncoder.SetIndent("", "\t")
	_ = jsonEncoder.Encode(data)
	log.Printf("%s", bf.String())
}

func UrlToBase64QrCode(url string) string {
	var png []byte
	png, _ = qrcode.Encode(url, qrcode.Medium, 256)
	base64Str := base64.StdEncoding.EncodeToString(png)
	return fmt.Sprintf("data:image/png;base64,%s", base64Str)
}

func JSON(v interface{}) ([]byte, error) {
	bf := bytes.NewBuffer([]byte{})
	jsonEncoder := json.NewEncoder(bf)
	jsonEncoder.SetEscapeHTML(false)
	err := jsonEncoder.Encode(v)
	if err != nil {
		return []byte{}, err
	}
	// 去除末尾的 \n
	return bf.Bytes()[:bf.Len()-1], nil
}

func JSONString(v interface{}) string {
	b, _ := JSON(v)
	return string(b)
}

func JSONParseString(d string, v interface{}) error {
	return json.Unmarshal([]byte(d), &v)
}
