package middleware_test

import (
	"os"
	"testing"

	"go-framework/pkg/boot"
)

var booted *boot.Booted

func TestMain(m *testing.M) {
	var err error
	if booted, err = boot.Boot(
		boot.WithConfigFile(os.Getenv("LGO_TEST_FILE")),
		boot.WithRoutePrint(false),
	); err != nil {
		panic(err)
	}
	m.Run()
}
