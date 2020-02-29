package creational

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstance(t *testing.T) {
	testInstance := GetInstance()
	assert.Equal(t, int64(0), testInstance.count)
	testInstance.Increment()
	assert.Equal(t, int64(1), testInstance.count)
}
