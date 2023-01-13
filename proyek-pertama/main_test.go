package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {
	text := HelloWorld()
	if text != "Hello World!" {
		t.Fail()
	}
}

func TestHelloWorld2(t *testing.T) {
	text := HelloWorld()
	assert.Equal(t, text, "Hello World!")
}
