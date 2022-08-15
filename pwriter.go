package pcollection

import "sync"

type PersistentWriter struct {
	Empty       bool         `json:"empty"`
	FilePath    string       `json:"filePath"`
	DataChan    chan any     `json:"dataChan"`
	ErrorsQueue *ErrorsQueue `json:"errorsQueue"`
	Once        *sync.Once   `json:"once"`
}

func NewPersistentWriter(filepath string, maxbuffersize int) (*PersistentWriter, error) {
	self := PersistentWriter{}
	self.Empty = filepath == ""
	self.FilePath = filepath
	self.DataChan = make(chan any, maxbuffersize)
	self.Once = new(sync.Once)
	self.ErrorsQueue = NewErrorsQueue(maxbuffersize)
	return &self, nil
}

func (pw *PersistentWriter) IsEmpty() bool {
	return pw.Empty
}
