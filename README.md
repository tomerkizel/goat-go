<div align="center">
  <img src="./Goat-logo.png" alt="Goat" title="Goat" height="300px" />
</div>

<div align="center">

# :goat:Persistent Data Structures For The Go Language:goat:

</div>

# Table of Contents
 - [Why Goat](#why-goat)
 - [Import Goat](#import-goat)
 - [PMap](#pmap)
	- [PMap Methods](#methods)
		- [AddOne and AddBatch](#addone-and-addbatch)
		- [Read](#read)
		- [Delete](#delete)
- [PArray](#parray)

# Why Goat?
Goat intoduces persistent data structures to Go.
Currently available types:
| Name | Implements | Section |
| ---- | ---- | ---- |
| PMap | Map | [PMap](#pmap) |
| PArray | Array | [PArray](#parray) | 

# Import Goat

```go
import github.com/tomerkizel/goat-go
```

# PMap

PMap is goat's map type implementation. To initizalize a PMap, use:

```go
self := goat.EmptyPMap(1, "")
```

In the above instance, 'self' will be a persistent map[int]string type.
To initizalize a persistent map[any]any type use:

```go
self := goat.EmptyPMap(nil, nil)
```

## Methods

### AddOne and AddBatch

```go
func (p *PMap) AddOne(key, value any) (*PMap, error)
func (p *PMap) AddBatch(keyvalue map[any]any) (*PMap, error)
```

AddOne adds a single key-value pair, and AddBatch adds a batch of key-value pairs to the persistent map. <b>key and value params must be of the right types.</b>
The methods returns a new PMap with the change/s

### Read

```go
func (p *PMap) Read(key any) (any, error)
```

Read returns the value of a given key in the PMap. the method works exaclty like trying to read a value of a regular map

### Delete

```go
func (p *PMap) Delete(key any) (*PMap, error) 
```

Delete removes a specific key-value pair from the PMap.
The method returns a new PMap with the change

## PArray

PArray is goat's array type implementation. To initizalize a PArray, use:

```go
self := EmptyPArray(1)
```

In the above instance, 'self' will be a persistent []int type.
To initizalize a persistent []any type use:

```go
self := goat.EmptyPArray(nil)
```