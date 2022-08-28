package goat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPMapAssignment(t *testing.T) {
	self := EmptyPMap(1, "hi")
	new, e := self.Set(2, "Da")
	assert.NoError(t, e)
	assert.NotEqual(t, new.GetMap(), self.GetMap())
	fail, e := self.Set("yo", 1)
	assert.Error(t, e)
	assert.Nil(t, fail)
	assert.Equal(t, new.Keys(), []any{2})
}

type TestStruct struct {
	IntVar  int            `json:"intVar"`
	BoolVar bool           `json:"boolVar"`
	MapVar  map[string]any `json:"mapVar"`
}

func TestPMapStructAssignment(t *testing.T) {
	self := EmptyPMap("", TestStruct{})
	new, e := self.Set("test", TestStruct{1, true, make(map[string]any, 5)})
	assert.NoError(t, e)
	assert.NotEqual(t, self.GetMap(), new.GetMap())
	fail, e := new.Set("fail", 1)
	assert.Error(t, e)
	assert.Nil(t, fail)
}

func TestPMapReAssignment(t *testing.T) {
	self := EmptyPMap(1, "")
	new, e := self.Set(1, "aaa")
	assert.NoError(t, e)
	assert.NotEqual(t, self.GetMap(), new.GetMap())
	val, e := new.Get(1)
	assert.NoError(t, e)
	assert.Equal(t, val, "aaa")
	assert.Equal(t, new.mapValue[1], "aaa")
	newer, e := new.Set(1, "bbb")
	assert.NoError(t, e)
	assert.NotEqual(t, new.GetMap(), newer.GetMap())
	val, e = newer.Get(1)
	assert.NoError(t, e)
	assert.Equal(t, val, "bbb")
}

func TestPMapMethods(t *testing.T) {
	arr := [4]string{"a", "b", "c", "d"}
	mapval := map[any]any{1: "a", 2: "b", 3: "c", 4: "d"}
	self := EmptyPMap(1, "")
	var e error
	for k, v := range mapval {
		self, e = self.Set(k, v)
		assert.NoError(t, e)
	}
	assert.Equal(t, self.GetMap(), mapval)
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
	new, e := self.Set(1, "hi")
	assert.NoError(t, e)
	assert.Equal(t, new.mapValue, map[any]any{1: "hi"})
	new, e = new.Set("hello", 2)
	assert.NoError(t, e)
	assert.Equal(t, new.mapValue, map[any]any{1: "hi", "hello": 2})
}

func TestPMapMerge(t *testing.T) {
	self := EmptyPMap(1, "")
	mapval := map[any]any{1: "a", 2: "b", 3: "c", 4: "d", 5: "e"}
	merger := EmptyPMap(1, "")
	ch := make(chan *PMap)
	var e error
	for k, v := range mapval {
		go func(v any) {
			new, e := self.Set(k, v)
			assert.NoError(t, e)
			ch <- new

		}(v)
		merger, e = merger.Merge(<-ch)
		assert.NoError(t, e)

	}
	assert.Equal(t, mapval, merger.GetMap())
}
