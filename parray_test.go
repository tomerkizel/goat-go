package goat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPArrayAssignment(t *testing.T) {
	self := EmptyPArray(1)
	new, e := self.AddOne(1)
	assert.NoError(t, e)
	assert.NotEqual(t, new.arrayValue, self.arrayValue)
	assert.Equal(t, new.arrayValue[0], 1)
	fail, e := self.AddOne("yo")
	assert.Error(t, e)
	assert.Nil(t, fail)
	arr := []any{1, 2, 3, 4, 5}
	newer, e := self.AddBatch(arr)
	assert.NoError(t, e)
	assert.Equal(t, newer.arrayValue, arr)
}

func TestPArrayStruct(t *testing.T) {
	self := EmptyPArray(TestStruct{})
	new, e := self.AddOne(TestStruct{1, true, make(map[string]any, 5)})
	assert.NoError(t, e)
	assert.NotEqual(t, self.arrayValue, new.arrayValue)
	fail, e := new.AddOne(1)
	assert.Error(t, e)
	assert.Nil(t, fail)
}

func TestPArrayReadDelete(t *testing.T) {
	self := EmptyPArray(1)
	arr := []any{1, 2, 3, 4, 5}
	new, e := self.AddBatch(arr)
	assert.NoError(t, e)
	val, e := new.Read(2)
	assert.Equal(t, val, 3)
	assert.NoError(t, e)
	_, e = new.Read(60)
	assert.Error(t, e)
	newer, e := new.Delete(1)
	assert.NoError(t, e)
	assert.NotEqual(t, newer.arrayValue, new.arrayValue)
	fail, e := newer.Delete(60)
	assert.Nil(t, fail)
	assert.Error(t, e)
}
