package pcollection

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/tomerkizel/goat-go/utils"

	"github.com/stretchr/testify/assert"
)

const (
	searchResult      = "SearchResult.json"
	emptySearchResult = "EmptySearchResult.json"
	unsortedFile      = "UnsortedFile.json"
	sortedFile        = "SortedFile.json"
	defaultKey        = "results"
)

type ArrayValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type inputRecord struct {
	IntKey   int          `json:"intKey"`
	StrKey   string       `json:"strKey"`
	BoolKey  bool         `json:"boolKey"`
	ArrayKey []ArrayValue `json:"arrayKey"`
}

func getTestDataPath() string {
	dir, _ := os.Getwd()
	return filepath.Join(dir, "testdata")
}

func TestPersistentCollectionConstructor(t *testing.T) {
	readPath := filepath.Join(getTestDataPath(), searchResult)
	pc, err := NewPersistentCollection(readPath, false, defaultKey, "")
	assert.Equal(t, err, nil)
	assert.Equal(t, pc.GetBufferSize(), 50000)
	assert.Equal(t, pc.GetReaderKey(), defaultKey)
}

func TestPersistentCollectionNextAndLength(t *testing.T) {
	readPath := filepath.Join(getTestDataPath(), searchResult)
	pc, err := NewPersistentCollection(readPath, false, defaultKey, "")
	assert.NoError(t, err)
	for i := 0; i < 2; i++ {

		var rSlice []inputRecord
		for item := new(inputRecord); pc.Next(item) == nil; item = new(inputRecord) {
			rSlice = append(rSlice, *item)
		}
		// First element
		assert.Equal(t, 1, rSlice[0].IntKey)
		assert.Equal(t, "A", rSlice[0].StrKey)
		assert.Equal(t, true, rSlice[0].BoolKey)
		assert.ElementsMatch(t, rSlice[0].ArrayKey, []ArrayValue{{Key: "build.number", Value: "6"}})
		// Second element
		assert.Equal(t, 2, rSlice[1].IntKey)
		assert.Equal(t, "B", rSlice[1].StrKey)
		assert.Equal(t, false, rSlice[1].BoolKey)
		assert.Empty(t, rSlice[1].ArrayKey)
		// Length validation
		length, err := pc.ReaderLength()
		assert.NoError(t, err)
		assert.Equal(t, 2, length)
		pc.ResetReader()
	}
}

func TestPersistentCollectionCloseReader(t *testing.T) {
	// Create a file.
	fd, err := utils.CreateTempFile()
	assert.NoError(t, err)
	assert.NoError(t, fd.Close())
	filePathToBeDeleted := fd.Name()

	// Load file to reader
	pc, err := NewPersistentCollection(filePathToBeDeleted, false, defaultKey, "")
	assert.NoError(t, err)
	// Check file exists
	_, err = os.Stat(filePathToBeDeleted)
	assert.NoError(t, err)

	// Check if the file got deleted
	assert.NoError(t, pc.Close())
	_, err = os.Stat(filePathToBeDeleted)
	assert.True(t, os.IsNotExist(err))
}
