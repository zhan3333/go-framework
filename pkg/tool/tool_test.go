package tool_test

import (
	"github.com/stretchr/testify/assert"
	"go-framework/pkg/tool"
	"testing"
)

func TestUrlToBase64QrCode(t *testing.T) {
	base64Str := tool.UrlToBase64QrCode("http://www.baidu.com")
	assert.NotEmpty(t, base64Str)
}
