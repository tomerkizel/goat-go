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
	Empty       bool                `json:"empty"`
	Length      int                 `json:"length"`
	FilePath    string              `json:"filePath"`
	JsonKey     string              `json:"jsonKey"`
	DataChan    chan map[string]any `json:"dataChan"`
	ErrorsQueue *ErrorsQueue        `json:"errorsQueue"`
	Once        *sync.Once          `json:"once"`
}

func NewPersistentReader(filepath, jsonkey string, maxbuffersize int) (*PersistentReader, error) {
	self := PersistentReader{}
	self.Empty = filepath == ""
	self.FilePath = filepath
	self.JsonKey = jsonkey
	self.DataChan = make(chan map[string]any, maxbuffersize)
	self.Once = new(sync.Once)
	self.ErrorsQueue = NewErrorsQueue(maxbuffersize)
	return &self, nil
}

func (pr *PersistentReader) IsEmpty() bool {
	return pr.Empty
}

func (pr *PersistentReader) NextRecord(recordOutput any) error {
	if pr.Empty {
		return fmt.Errorf("Empty")
	}
	pr.Once.Do(func() {
		go func() {
			defer close(pr.DataChan)
			pr.Length = 0
			pr.readSingleFile()
		}()
	})
	record, ok := <-pr.DataChan
	if !ok {
		return io.EOF
	}
	bytes, err := json.Marshal(record)
	if err != nil {
		pr.ErrorsQueue.Add(err)
		return err
	}
	err = json.Unmarshal(bytes, &recordOutput)
	if err != nil {
		pr.ErrorsQueue.Add(err)
		return err
	}
	pr.Length++
	return err
}

func (pr *PersistentReader) readSingleFile() {
	fd, err := os.Open(pr.FilePath)
	if err != nil {
		pr.ErrorsQueue.Add(err)
		return
	}
	defer func() {
		err = fd.Close()
		if err != nil {
			pr.ErrorsQueue.Add(err)
		}
	}()
	br := bufio.NewReaderSize(fd, 65536) //65536 from original code, what does it mean?
	dec := json.NewDecoder(br)
	err = utils.FindDecoderTargetPosition(dec, pr.JsonKey)
	if err != nil {
		if err == io.EOF {
			pr.ErrorsQueue.Add(fmt.Errorf((pr.JsonKey + " not found")))
			return
		}
		pr.ErrorsQueue.Add(err)
		return
	}
	for dec.More() {
		var ResultItem map[string]any
		err := dec.Decode(&ResultItem)
		if err != nil {
			pr.ErrorsQueue.Add(err)
			return
		}
		pr.DataChan <- ResultItem
	}
}

// Prepare the reader to read the file all over again (not thread-safe).
func (pr *PersistentReader) Reset() {
	pr.DataChan = make(chan map[string]any, cap(pr.DataChan))
	pr.Once = new(sync.Once)
}

// Cleanup the reader data.
func (pr *PersistentReader) Close() error {
	if pr.FilePath == "" {
		return nil
	}
	if err := os.Remove(pr.FilePath); err != nil {
		return errors.New("Failed to close reader: " + err.Error())
	}

	pr.FilePath = ""
	return nil
}
