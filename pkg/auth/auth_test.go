package auth_test

import (
	"github.com/stretchr/testify/assert"
	"go-framework/pkg/auth"
	"testing"
)

func TestEncrypt(t *testing.T) {
	password := "123456"
	enc, err := auth.Encrypt(password)
	assert.Nil(t, err)
	assert.NotEmpty(t, enc)
	assert.Nil(t, auth.Compare(enc, password))
}
