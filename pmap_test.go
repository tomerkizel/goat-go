package goat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAssignment(t *testing.T) {
	self := EmptyPMap(1, "hi")
	new, e := self.Add(2, "Da")
	assert.NoError(t, e)
	assert.NotEqual(t, new.mapValue, self.mapValue)
	fail, e := self.Add("yo", 1)
	assert.Error(t, e)
	assert.Nil(t, fail)
}

type TestStruct struct {
	IntVar  int            `json:"intVar"`
	BoolVar bool           `json:"boolVar"`
	MapVar  map[string]any `json:"mapVar"`
}

func TestStructAssignment(t *testing.T) {
	self := EmptyPMap("", TestStruct{})
	new, e := self.Add("test", TestStruct{1, true, make(map[string]any, 5)})
	assert.NoError(t, e)
	assert.NotEqual(t, self.mapValue, new.mapValue)
	fail, e := new.Add("fail", 1)
	assert.Error(t, e)
	assert.Nil(t, fail)
}

func TestReAssignment(t *testing.T) {
	self := EmptyPMap(1, "")
	new, e := self.Add(1, "aaa")
	assert.NoError(t, e)
	assert.NotEqual(t, self.mapValue, new.mapValue)
	assert.Equal(t, new.mapValue[1], "aaa")
	newer, e := new.Add(1, "bbb")
	assert.NoError(t, e)
	assert.NotEqual(t, new.mapValue, newer.mapValue)
	assert.Equal(t, newer.mapValue[1], "bbb")
}
