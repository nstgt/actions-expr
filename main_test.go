package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_say(t *testing.T) {
	s := "I'm Gopher!"
	expected := "ʕ◔ϖ◔ʔ > I'm Gopher!"
	assert.Equal(t, gopherSay(s), expected)
}
