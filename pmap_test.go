package pmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCorrectAssignment(t *testing.T) {
	self := Empty(1, "hi")
	new, e := self.Add(2, "Da")
	assert.NoError(t, e)
	assert.NotEqual(t, new.mapValue, self.mapValue)
}

func TestWrongAssignment(t *testing.T) {
	self := Empty(1, "hi")
	new, e := self.Add("yo", 1)
	assert.Error(t, e)
	assert.Nil(t, new)
}
