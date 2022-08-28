package goat

import (
	"sync"
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
	assert.Equal(t, self.arrayValue, arr)
}

func TestPArrayStruct(t *testing.T) {
	self := EmptyPArray(TestStruct{})
	new, e := self.Push(TestStruct{1, true, make(map[string]any, 5)})
	assert.NoError(t, e)
	assert.NotEqual(t, self.arrayValue, new.arrayValue)
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
	assert.NotEqual(t, self.arrayValue, new_1.arrayValue)
	assert.Equal(t, item, 2)
	fail, item, e := self.Delete(60)
	assert.Nil(t, fail)
	assert.Nil(t, item)
	assert.Error(t, e)
	new_2, item, e := self.Pop()
	assert.NoError(t, e)
	assert.Equal(t, item, 5)
	assert.NotEqual(t, self.arrayValue, new_2.arrayValue)
	fail, none, e := EmptyPArray(1).Pop()
	assert.Error(t, e)
	assert.Nil(t, none)
	assert.Nil(t, fail)
	new_3, e := self.Set(1, 60)
	assert.NoError(t, e)
	assert.Equal(t, new_3.arrayValue, []any{1, 60, 3, 4, 5})
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
	assert.Equal(t, new.arrayValue, []any{1})
	new, e = new.Push("hi")
	assert.NoError(t, e)
	assert.Equal(t, new.arrayValue, []any{1, "hi"})
}

func TestPArrayMergeAndSort(t *testing.T) {
	self := EmptyPArray(1)
	arr := []any{1, 2, 3, 4, 5}
	merger := EmptyPArray(1)
	var wait sync.WaitGroup
	for _, v := range arr {
		wait.Add(1)
		go func(v any) {
			defer wait.Done()
			new, e := self.Push(v)
			assert.NoError(t, e)
			merger, e = merger.Merge(new)
			assert.NoError(t, e)
		}(v)
	}
	wait.Wait()
	x := func(i, j any) bool {
		item_x, ok := i.(int)
		assert.Equal(t, ok, true)
		item_y, ok := j.(int)
		assert.Equal(t, ok, true)
		return item_x < item_y
	}
	sorted, e := merger.Sort(x)
	assert.NoError(t, e)
	val := sorted.GetArray()
	assert.Equal(t, val, arr)
}
