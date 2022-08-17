package goat

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/tomerkizel/goat-go/utils"
)

type PersistentReader struct {
	empty       bool
	length      int
	filePath    string
	jsonKey     string
	dataChan    chan map[string]any
	errorsQueue *ErrorsQueue
	once        *sync.Once
}

func newPersistentReader(filepath, jsonkey string, maxbuffersize int) (*PersistentReader, error) {
	self := PersistentReader{}
	self.empty = filepath == ""
	self.filePath = filepath
	self.jsonKey = jsonkey
	self.dataChan = make(chan map[string]any, maxbuffersize)
	self.once = new(sync.Once)
	self.errorsQueue = NewErrorsQueue(maxbuffersize)
	return &self, nil
}

func (pr *PersistentReader) IsEmpty() bool {
	return pr.empty
}

func (pr *PersistentReader) GetJsonKey() string {
	return pr.jsonKey
}

func (pr *PersistentReader) nextRecord(recordOutput any) error {
	if pr.IsEmpty() {
		return fmt.Errorf("Empty")
	}
	pr.once.Do(func() {
		go func() {
			defer close(pr.dataChan)
			pr.length = 0
			pr.readSingleFile()
		}()
	})
	record, ok := <-pr.dataChan
	if !ok {
		return io.EOF
	}
	bytes, err := json.Marshal(record)
	if err != nil {
		pr.errorsQueue.Add(err)
		return err
	}
	err = json.Unmarshal(bytes, &recordOutput)
	if err != nil {
		pr.errorsQueue.Add(err)
		return err
	}
	pr.length++
	return err
}

func (pr *PersistentReader) readSingleFile() {
	fd, err := os.Open(pr.filePath)
	if err != nil {
		pr.errorsQueue.Add(err)
		return
	}
	defer func() {
		err = fd.Close()
		if err != nil {
			pr.errorsQueue.Add(err)
		}
	}()
	br := bufio.NewReaderSize(fd, 65536) //65536 from original code, what does it mean?
	dec := json.NewDecoder(br)
	err = utils.FindDecoderTargetPosition(dec, pr.jsonKey)
	if err != nil {
		if err == io.EOF {
			pr.errorsQueue.Add(fmt.Errorf((pr.jsonKey + " not found")))
			return
		}
		pr.errorsQueue.Add(err)
		return
	}
	for dec.More() {
		var ResultItem map[string]any
		err := dec.Decode(&ResultItem)
		if err != nil {
			pr.errorsQueue.Add(err)
			return
		}
		pr.dataChan <- ResultItem
	}
}

// Prepare the reader to read the file all over again (not thread-safe).
func (pr *PersistentReader) reset() {
	pr.dataChan = make(chan map[string]any, cap(pr.dataChan))
	pr.once = new(sync.Once)
}

// Cleanup the reader data.
func (pr *PersistentReader) close() error {
	if pr.filePath == "" {
		return nil
	}
	if err := os.Remove(pr.filePath); err != nil {
		return errors.New("Failed to close reader: " + err.Error())
	}

	pr.filePath = ""
	return nil
}

func (pr *PersistentReader) getError() error {
	return pr.errorsQueue.Get()
}
