package goat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPArrayAssignment(t *testing.T) {
	self := EmptyPArray(1)
	new, e := self.AddOne(1)
	assert.NoError(t, e)
	assert.NotEqual(t, new.valueArr, self.valueArr)
	assert.Equal(t, new.valueArr[0], 1)
	fail, e := self.AddOne("yo")
	assert.Error(t, e)
	assert.Nil(t, fail)
}
