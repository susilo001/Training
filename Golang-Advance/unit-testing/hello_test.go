package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloFunction(t *testing.T) {
	message := GetHelloMessage()

	assert.Equal(t, "Hello Hello", message)
}
