package goat

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
)

type PersistentWriter struct {
	empty        bool
	completeFile bool
	outputFile   *os.File
	dataChan     chan any
	errorsQueue  *ErrorsQueue
	once         *sync.Once
	jsonKey      string
	runWait      sync.WaitGroup
}

func newPersistentWriter(jsonkey string, iscompletefile bool, maxbuffersize int) (*PersistentWriter, error) {
	self := PersistentWriter{}
	self.completeFile = iscompletefile
	self.empty = true
	self.dataChan = make(chan any, maxbuffersize)
	self.once = new(sync.Once)
	self.errorsQueue = NewErrorsQueue(maxbuffersize)
	self.jsonKey = jsonkey
	return &self, nil
}

func (pw *PersistentWriter) IsEmpty() bool {
	return pw.empty
}

func (pw *PersistentWriter) IsCompleteFile() bool {
	return pw.completeFile
}

func (pw *PersistentWriter) GetFilePath() string {
	if pw.outputFile != nil {
		return pw.outputFile.Name()
	}
	return ""
}

func (pw *PersistentWriter) GetJsonKey() string {
	return pw.jsonKey
}

func (pw *PersistentWriter) write(record any) {
	pw.empty = false
	pw.startWritingWorker()
	pw.dataChan <- record
}

func (pw *PersistentWriter) startWritingWorker() {
	pw.once.Do(func() {
		var err error
		pw.outputFile, err = utils.CreateTempFile()
		if err != nil {
			pw.errorsQueue.Add(err)
			return
		}
		pw.runWait.Add(1)
		go func() {
			defer pw.runWait.Done()
			pw.run()
		}()
	})
}

func (pw *PersistentWriter) run() {
	var err error
	defer func() {
		err = pw.outputFile.Close()
		if err != nil {
			pw.errorsQueue.Add(err)
		}
	}()
	openString := jsonArrayPrefixPattern
	closeString := ""
	if pw.IsCompleteFile() {
		openString = "{\n" + openString
	}
	_, err = pw.outputFile.WriteString(fmt.Sprintf(openString, pw.GetJsonKey()))
	if err != nil {
		pw.errorsQueue.Add(err)
		return
	}
	buf := bytes.NewBuffer(nil)
	enc := json.NewEncoder(buf)
	enc.SetIndent("    ", "  ")
	recordPrefix := "\n    "
	firstRecord := true
	for record := range pw.dataChan {
		buf.Reset()
		err = enc.Encode(record)
		if err != nil {
			pw.errorsQueue.Add(err)
			continue
		}
		record := recordPrefix + string(bytes.TrimRight(buf.Bytes(), "\n"))
		_, err = pw.outputFile.WriteString(record)
		if err != nil {
			pw.errorsQueue.Add(err)
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
	_, err = pw.outputFile.WriteString(closeString)
	if err != nil {
		pw.errorsQueue.Add(err)
	}
}

func (pw *PersistentWriter) getError() error {
	return pw.errorsQueue.Get()
}

func (pw *PersistentWriter) removeOutputFilePath() error {
	return os.Remove(pw.outputFile.Name())
}

func (pw *PersistentWriter) close() error {
	if pw.IsEmpty() {
		return nil
	}
	close(pw.dataChan)
	pw.runWait.Wait()
	if err := pw.getError(); err != nil {
		pw.errorsQueue.Add(err)
		return err
	}
	return nil
}
