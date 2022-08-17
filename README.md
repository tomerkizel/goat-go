<dic align="center">
  <img src="./Goat-logo.png" alt="Goat" title="Goat" height="300px" />
</div>

<div align="center">

# :goat:Persistent Collection For The Go Language:goat:

</div>

---

<div align="center">
	This package was first written on and extracted from <a href="https://github.com/jfrog/jfrog-client-go">JFrog client go</a>
</div>

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

To initialize a new persistent collection, use the goat.NewPersistentCollection function:
```go
pc, err := goat.NewPersistentCollection(readPath, false, "results", "")
```
The parameters send to NewPersistentCollection are:
| Name | Type | Info |
| ---- | ---- | ---- |
| readfilepath | string | The JSON file path from which the collection will read |
| writefullfile | bool | Flag if the collection will write a full JSON file |
| readkey | string | The key of the JSON object the collection will read - must be of type array |
| writekey | string | 28The key of the JSON object the collection will write3 |

After initializing a new persistent collection, you can read chunks from the file by using the Next method:
```go
pc, err := goat.NewPersistentCollection(readPath, false, "results", "")	
var output []inputRecord
for item := new(inputRecord); pc.Next(item) == nil; item = new(inputRecord) {
	output = append(rSlice, *item)
}
```
output will include all the elements of the array read from the JSON located readPath in key "results"

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
	Empty        bool           `json:"empty"`
	CompleteFile bool           `json:"completeFile"`
	OutputFile   *os.File       `json:"outputFile"`
	DataChan     chan any       `json:"dataChan"`
	ErrorsQueue  *ErrorsQueue   `json:"errorsQueue"`
	Once         *sync.Once     `json:"once"`
	JsonKey      string         `json:"jsonKey"`
	RunWait      sync.WaitGroup `json:"runWait"`
}
```