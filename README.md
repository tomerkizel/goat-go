<h1 align="center">
  <img src="./Goat-logo.png" alt="Goat" title="Goat" height="300px" />
</h1>

<h2 align="center"> Goat </h2>
<h2 align="center">Persistent Collection For The Go Language </h2>

<h7 align="center"> This package was first written on <a href="https://github.com/jfrog/jfrog-client-go">JFrog client go</a></h3>

---

## Table of Contents
 - [Why Goat](#why-goat)
 - [Import Goat](#import-goat)
 - [Using Goat](#using-goat)
    - [PersistentCollection](#persistentcollection)
    - [PersistentReader](#persistentreader)
    - [PersistentWriter](#persistentwriter)

## Why Goat?
Goat allows you to create a persistent collection, that reads and writes JSON files using small chunks of memory.
The default chunk size is 50000 but it can be change, see [Using Goat](#using-goat)

## Import Goat

```go
    import github.com/tomerkizel/goat-go
```

## Using Goat

### PersistentCollection
PersistentCollection is the main struct you'll be using while working with Goat.

```go
type PersistentCollection struct {
	MaxBufferSize    int               `json:"maxBufferSize"`
	PersistentReader *PersistentReader `json:"persistentReader"`
	PersistentWriter *PersistentWriter `json:"persistentWriter"`
}
```

### PersistentReader
PersistentReader is used to read JSON files into small chunks of memory

```go
type PersistentReader struct {
	Empty       bool                `json:"empty"`
	Length      int                 `json:"length"`
	FilePath    string              `json:"filePath"`
	JsonKey     string              `json:"jsonKey"`
	DataChan    chan map[string]any `json:"dataChan"`
	ErrorsQueue *ErrorsQueue        `json:"errorsQueue"`
	Once        *sync.Once          `json:"once"`
}
```

### PersistentWriter
PersistentWriter is used to write JSON files from small chunks of memory

```go
type PersistentWriter struct {
	Empty       bool         `json:"empty"`
	FilePath    string       `json:"filePath"`
	DataChan    chan any     `json:"dataChan"`
	ErrorsQueue *ErrorsQueue `json:"errorsQueue"`
	Once        *sync.Once   `json:"once"`
}
```