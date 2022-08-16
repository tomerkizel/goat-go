package pcollection

import (
	"os"
	"sync"
)

type PersistentWriter struct {
	Empty       bool         `json:"empty"`
	OutputFile  *os.File     `json:"outputFile"`
	DataChan    chan any     `json:"dataChan"`
	ErrorsQueue *ErrorsQueue `json:"errorsQueue"`
	Once        *sync.Once   `json:"once"`
	JsonKey     string       `json:"jsonKey"`
}

func NewPersistentWriter(jsonkey string, outputfile *os.File, maxbuffersize int) (*PersistentWriter, error) {
	self := PersistentWriter{}
	self.Empty = outputfile == nil
	self.OutputFile = outputfile
	self.DataChan = make(chan any, maxbuffersize)
	self.Once = new(sync.Once)
	self.ErrorsQueue = NewErrorsQueue(maxbuffersize)
	self.JsonKey = jsonkey
	return &self, nil
}

func (pw *PersistentWriter) IsEmpty() bool {
	return pw.Empty
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
