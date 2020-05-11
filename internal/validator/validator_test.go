package validator_test

import (
	"github.com/stretchr/testify/assert"
	"go-framework/bootstrap"
	"go-framework/internal/validator"
	v "gopkg.in/go-playground/validator.v9"
	"testing"
)

func TestMain(m *testing.M) {
	bootstrap.SetInTest()
	bootstrap.Bootstrap()
	m.Run()
}

func TestMobileValidator(t *testing.T) {
	var err error
	type user struct {
		Mobile string `json:"mobile" validate:"required,mobile"`
	}
	u := user{}
	validate := v.New()
	_ = validate.RegisterValidation("mobile", validator.Mobile)
	err = validate.Struct(u)
	assert.NotNil(t, err)
	if err != nil {
		t.Logf("err: %s", err.Error())
	}
	u.Mobile = "13517210601"
	err = validate.Struct(u)
	assert.Nil(t, err)
}
