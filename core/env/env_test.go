package env_test

import (
	"github.com/stretchr/testify/assert"
	"go-framework/core/env"
	"os"
	"testing"
)

func TestDefaultGet(t *testing.T) {
	assert.Equal(t, os.Getenv("APP_NAME"), env.DefaultGet("APP_NAME", ""))
}
