package goat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPMapAssignment(t *testing.T) {
	self := EmptyPMap(1, "hi")
	new, e := self.SetOne(2, "Da")
	assert.NoError(t, e)
	assert.NotEqual(t, new.mapValue, self.mapValue)
	fail, e := self.SetOne("yo", 1)
	assert.Error(t, e)
	assert.Nil(t, fail)
	keys := new.Keys()
	assert.Equal(t, keys, []any{2})
}

type TestStruct struct {
	IntVar  int            `json:"intVar"`
	BoolVar bool           `json:"boolVar"`
	MapVar  map[string]any `json:"mapVar"`
}

func TestPMapStructAssignment(t *testing.T) {
	self := EmptyPMap("", TestStruct{})
	new, e := self.SetOne("test", TestStruct{1, true, make(map[string]any, 5)})
	assert.NoError(t, e)
	assert.NotEqual(t, self.mapValue, new.mapValue)
	fail, e := new.SetOne("fail", 1)
	assert.Error(t, e)
	assert.Nil(t, fail)
}

func TestPMapReAssignment(t *testing.T) {
	self := EmptyPMap(1, "")
	new, e := self.SetOne(1, "aaa")
	assert.NoError(t, e)
	assert.NotEqual(t, self.mapValue, new.mapValue)
	assert.Equal(t, new.mapValue[1], "aaa")
	newer, e := new.SetOne(1, "bbb")
	assert.NoError(t, e)
	assert.NotEqual(t, new.mapValue, newer.mapValue)
	assert.Equal(t, newer.mapValue[1], "bbb")
}

func TestPMapMethods(t *testing.T) {
	arr := [4]string{"a", "b", "c", "d"}
	mapval := map[any]any{1: "a", 2: "b", 3: "c", 4: "d"}
	self := EmptyPMap(1, "")
	self, e := self.SetMany(mapval)
	assert.NoError(t, e)
	for i := 1; i <= 4; i++ {
		val, e := self.Get(i)
		assert.NoError(t, e)
		assert.Equal(t, val, arr[i-1])
	}
	val, e := self.Get("1")
	assert.Error(t, e)
	assert.Nil(t, val)
	fail, item, e := self.Delete("ar")
	assert.Error(t, e)
	assert.Nil(t, fail)
	assert.Nil(t, item)
	suc, item, e := self.Delete(1)
	assert.NoError(t, e)
	assert.Nil(t, suc.mapValue[1])
	assert.Equal(t, item, "a")
}

func TestPMapNil(t *testing.T) {
	self := EmptyPMap(nil, nil)
	new, e := self.SetOne(1, "hi")
	assert.NoError(t, e)
	assert.Equal(t, new.mapValue, map[any]any{1: "hi"})
	new, e = new.SetOne("hello", 2)
	assert.NoError(t, e)
	assert.Equal(t, new.mapValue, map[any]any{1: "hi", "hello": 2})
}
