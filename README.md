<div align="center">
  <img src="./Goat-logo.png" alt="Goat" title="Goat" height="300px" />
</div>

<div align="center">

# :goat:Persistent Data Structures For The Go Language:goat:

</div>

# Why Goat?
Goat intoduces persistent data structures to Go. Persistent data struct is a data structure that always preserves the previous version of itself when it is modified. Therefor, each modifing method returns a new instance of the data structure with the modifications while keeping the original instance unchanged.

Currently available types:

| Name | Implements | Section |
| ---- | ---- | ---- |
| PMap | Map | [PMap](#pmap) |
| PArray | Array | [PArray](#parray) | 

# Table of Contents
 - [Import Goat](#import-goat)
 - [PMap](#pmap)
	- [PMap Methods](#key-methods)
	- [PMap Examples](#examples)
- [PArray](#parray)
	- [PArray Methods](#methods-1)
	- [PArray Examples](#examples-1)

# Import Goat

```go
import github.com/tomerkizel/goat-go
```

# PMap

PMap is goat's map type implementation. To initizalize a PMap, use:

```go
self := goat.EmptyPMap(x, y)
```
self will be a persistent map[typeof(x)]typeof(y) type. Pass nil, nil to generate a persistent map[any]any type.

## Key Methods
Delete removes a specific key-value pair from the PMap. The method returns a new PMap with the change, the removed value of the key and an error.
```go
func (p *PMap) Delete(key any) (*PMap, any, error)
```

Keys returns an array of the keys of the PMap
```go
func (p *PMap) Keys() []any
```

Merge allows you to merge two PMaps of the same key-value types into a single PMap
```go
func (p *PArray) Merge(q *PArray) (*PArray, error)
```

## Examples


# PArray

PArray is goat's array type implementation. To initizalize a PArray, use:

```go
self := goat.EmptyPArray(x)
```
self will be a persistent []typeof(x) type. Pass nil to generate a persistent []any type.


## Key Methods
Sort sorts the PArray with a given function
```go
func (p *PArray) Sort(fn func(x, y any) bool) (*PArray, error)
```
GetArray returns a copy of the array inside PArray
```go
func (p *PArray) GetArray() []any
```
Merge allows you to merge two PArrays of the same type into a single PArray
```go
func (p *PArray) Merge(q *PArray) (*PArray, error)
```
## Examples
The blow code generates an empty integer PArray, then, using PArray.Merge(..), the values of arr are being pushed into merger

(<b>Note:</b> merger.Merge(..) is assigned to merger, if this wasn't the case, merger would remain empty)

sorted will be a new PArray instance, sorted by x function. The code output will be <i>1 2 3 4 6 </i>
```go
self := goat.EmptyPArray(1)
arr := []any{1, 2, 3, 4, 5}
ch := make(chan *PArray)
var e error
for _, v := range arr {
	go func(v any) {
		new, e := self.Push(v)
		if e != nil {
			return
		}
		ch <- new
	}(v)
	merger, e = merger.Merge(<-ch)
	if e != nil {
		return
	}
}
x := func(i, j any) bool {
	item_x, ok := i.(int)
	if !ok {
		return false
	}
	item_y, ok := j.(int)
	if !ok {
		return false
	}
	return item_x < item_y
}
sorted, e := merger.Sort(x)
if e != nil {
	return
}
for _, v := range sorted.GetArray() {
	fmt.Print("%v ", v)
}
```
