package pcollection

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"sync"

	"github.com/tomerkizel/goat-go/utils"
)

const (
	jsonArrayPrefixPattern = "  \"%s\": ["
	jsonArraySuffix        = "]\n"
	DefaultKey             = "results"
)

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

func NewPersistentWriter(jsonkey string, iscompletefile bool, maxbuffersize int) (*PersistentWriter, error) {
	self := PersistentWriter{}
	self.CompleteFile = iscompletefile
	self.Empty = true
	self.DataChan = make(chan any, maxbuffersize)
	self.Once = new(sync.Once)
	self.ErrorsQueue = NewErrorsQueue(maxbuffersize)
	self.JsonKey = jsonkey
	return &self, nil
}

func (pw *PersistentWriter) IsEmpty() bool {
	return pw.Empty
}

func (pw *PersistentWriter) IsCompleteFile() bool {
	return pw.CompleteFile
}

func (pw *PersistentWriter) GetFilePath() string {
	if pw.OutputFile != nil {
		return pw.OutputFile.Name()
	}
	return ""
}

func (pw *PersistentWriter) GetJsonKey() string {
	return pw.JsonKey
}

func (pw *PersistentWriter) Write(record any) {
	pw.Empty = false
	pw.startWritingWorker()
	pw.DataChan <- record
}

func (pw *PersistentWriter) startWritingWorker() {
	pw.Once.Do(func() {
		var err error
		pw.OutputFile, err = utils.CreateTempFile()
		if err != nil {
			pw.ErrorsQueue.Add(err)
			return
		}
		pw.RunWait.Add(1)
		go func() {
			defer pw.RunWait.Done()
			pw.run()
		}()
	})
}

func (pw *PersistentWriter) run() {
	var err error
	defer func() {
		err = pw.OutputFile.Close()
		if err != nil {
			pw.ErrorsQueue.Add(err)
		}
	}()
	openString := jsonArrayPrefixPattern
	closeString := ""
	if pw.IsCompleteFile() {
		openString = "{\n" + openString
	}
	_, err = pw.OutputFile.WriteString(fmt.Sprintf(openString, pw.GetJsonKey()))
	if err != nil {
		pw.ErrorsQueue.Add(err)
		return
	}
	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)
	enc.SetIndent("    ", "  ")
	recordPrefix := "\n    "
	firstRecord := true
	for record := range pw.DataChan {
		buf.Reset()
		err = enc.Encode(record)
		if err != nil {
			pw.ErrorsQueue.Add(err)
			continue
		}
		record := recordPrefix + string(bytes.TrimRight(buf.Bytes(), "\n"))
		_, err = pw.OutputFile.WriteString(record)
		if err != nil {
			pw.ErrorsQueue.Add(err)
			continue
		}
		if firstRecord {
			// If a record was printed, we want to print a comma and ne before each and every future record.
			recordPrefix = "," + recordPrefix
			// We will close the array in a new-indent line.
			closeString = "\n  "
			firstRecord = false
		}
	}
	closeString = closeString + jsonArraySuffix
	if pw.IsCompleteFile() {
		closeString += "}\n"
	}
	_, err = pw.OutputFile.WriteString(closeString)
	if err != nil {
		pw.ErrorsQueue.Add(err)
	}
}

func (pw *PersistentWriter) Close() error {
	if pw.IsEmpty() {
		return nil
	}
	close(pw.DataChan)
	pw.RunWait.Wait()
	if err := pw.GetError(); err != nil {
		pw.ErrorsQueue.Add(err)
		return err
	}
	return nil
}

func (pw *PersistentWriter) GetError() error {
	return pw.ErrorsQueue.Get()
}

func (pw *PersistentWriter) RemoveOutputFilePath() error {
	return os.Remove(pw.OutputFile.Name())
}
