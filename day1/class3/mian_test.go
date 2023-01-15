package main

/**
单元测试
*/
import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func HelloTom() string {
	return "Tom"
}

func TestHelloTom(t *testing.T) {
	output := HelloTom()
	expectOutput := "Tom"
	assert.Equal(t, expectOutput, output)
}
