package validator_test

import (
	"github.com/stretchr/testify/assert"
	v "gopkg.in/go-playground/validator.v9"
	"os"
	"testing"

	"go-framework/internal/validator"
	"go-framework/pkg/boot"
)

func TestMain(m *testing.M) {
	if _, err := boot.Boot(
		boot.WithConfigFile(os.Getenv("LGO_TEST_FILE")),
		boot.WithRoutePrint(false),
	); err != nil {
		panic(err)
	}
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
		t.Logf("err: %v", err)
	}
	u.Mobile = "13517210601"
	err = validate.Struct(u)
	assert.Nil(t, err)
}
