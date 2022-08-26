<div align="center">
  <img src="./Goat-logo.png" alt="Goat" title="Goat" height="300px" />
</div>

<div align="center">

# :goat:Persistent Data Structures For The Go Language:goat:

</div>

## Table of Contents
 - [Why Goat](#why-goat)
 - [Import Goat](#import-goat)
 - [Using Goat](#using-goat)
    - [PMap](#pmap)

## Why Goat?
Goat intoduces persistent data structures to Go.
Currently available types:
| Name | Implements | Section |
| ---- | ---- | ---- |
| PMap | Map | [PMap](#pmap) |
| PArray | Array | [PArray](#parray) | 

## Import Goat

```go
import github.com/tomerkizel/goat-go
```

## Using Goat

### PMap

PMap is goat's map type implementation.

To initizalize a PMap, use:

```go
	self := goat.EmptyPMap(1, "")
```

In the above instance, 'self' will be a persistent map[int]string type
To initizalize a persistent map[any]any type use:

```go
	self := goat.EmptyPMap(nil, nil)
```