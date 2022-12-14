package goat

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPArrayAssignment(t *testing.T) {
	self := EmptyPArray(1)
	new, e := self.Push(1)
	assert.NoError(t, e)
	assert.NotEqual(t, new.arrayValue, self.arrayValue)
	assert.Equal(t, new.arrayValue[0], 1)
	fail, e := self.Push("yo")
	assert.Error(t, e)
	assert.Nil(t, fail)
	arr := []any{1, 2, 3, 4, 5}
	for _, v := range arr {
		self, e = self.Push(v)
		assert.NoError(t, e)
	}
	assert.Equal(t, self.GetArray(), arr)
}

func TestPArrayStruct(t *testing.T) {
	self := EmptyPArray(TestStruct{})
	new, e := self.Push(TestStruct{1, true, make(map[string]any, 5)})
	assert.NoError(t, e)
	assert.NotEqual(t, self.GetArray(), new.GetArray())
	fail, e := new.Push(1)
	assert.Error(t, e)
	assert.Nil(t, fail)
}

func TestPArrayMethods(t *testing.T) {
	arr := []any{1, 2, 3, 4, 5}
	self := EmptyPArray(1)
	var e error
	for _, v := range arr {
		self, e = self.Push(v)
		assert.NoError(t, e)
	}
	new_1, item, e := self.Delete(1)
	assert.NoError(t, e)
	assert.NotEqual(t, self.GetArray(), new_1.GetArray())
	assert.Equal(t, item, 2)
	fail, item, e := self.Delete(60)
	assert.Nil(t, fail)
	assert.Nil(t, item)
	assert.Error(t, e)
	new_2, item, e := self.Pop()
	assert.NoError(t, e)
	assert.Equal(t, item, 5)
	assert.NotEqual(t, self.GetArray(), new_2.GetArray())
	fail, none, e := EmptyPArray(1).Pop()
	assert.Error(t, e)
	assert.Nil(t, none)
	assert.Nil(t, fail)
	new_3, e := self.Set(1, 60)
	assert.NoError(t, e)
	assert.Equal(t, new_3.GetArray(), []any{1, 60, 3, 4, 5})
	fail, e = new_3.Set(2, "hi")
	assert.Error(t, e)
	assert.Nil(t, fail)
	fail, e = new_3.Set(60, 3)
	assert.Error(t, e)
	assert.Nil(t, fail)
}

func TestPArrayNil(t *testing.T) {
	self := EmptyPArray(nil)
	new, e := self.Push(1)
	assert.NoError(t, e)
	assert.Equal(t, new.GetArray(), []any{1})
	new, e = new.Push("hi")
	assert.NoError(t, e)
	assert.Equal(t, new.GetArray(), []any{1, "hi"})
}

func TestPArrayMergeAndSort(t *testing.T) {
	self := EmptyPArray(1)
	arr := []any{1, 2, 3, 4, 5}
	merger := EmptyPArray(1)
	ch := make(chan *PArray)
	var e error
	for _, v := range arr {
		go func(v any) {
			new, e := self.Push(v)
			assert.NoError(t, e)
			ch <- new

		}(v)
		merger, e = merger.Merge(<-ch)
		assert.NoError(t, e)

	}
	x := func(i, j any) bool {
		item_x, ok := i.(int)
		assert.Equal(t, ok, true)
		item_y, ok := j.(int)
		assert.Equal(t, ok, true)
		return item_x < item_y
	}
	sorted, e := merger.Sort(x)
	assert.NoError(t, e)
	assert.Equal(t, sorted.GetArray(), arr)
}
