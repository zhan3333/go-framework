package randc_test

import (
	"github.com/stretchr/testify/assert"
	"go-framework/pkg/randc"
	"testing"
)

func TestRandStringN(t *testing.T) {
	randStr := randc.RandStringN(10)
	assert.Equal(t, 10, len(randStr))
}
