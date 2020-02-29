package creational

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetInstance(t *testing.T) {
	testInstance := GetInstance()
	assert.Equal(t, testInstance.count, int64(0))
	testInstance.Increment()
	assert.Equal(t, testInstance.count, int64(1))
}
