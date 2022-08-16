package pcollection

import "os"

type PersistentCollection struct {
	MaxBufferSize    int               `json:"maxBufferSize"`
	PersistentReader *PersistentReader `json:"persistentReader"`
	PersistentWriter *PersistentWriter `json:"persistentWriter"`
}

func NewPersistentCollection(readfilepath string, writefilepath *os.File, readkey, writekey string) (*PersistentCollection, error) {
	self := NewEmptyPersistentCollection()
	err := self.SetReader(readkey, readfilepath)
	if err != nil {
		return nil, err
	}
	err = self.SetWriter(writekey, writefilepath)
	if err != nil {
		return nil, err
	}
	return self, nil
}

func NewEmptyPersistentCollection() *PersistentCollection {
	self := PersistentCollection{}
	self.MaxBufferSize = 50000
	return &self
}

func (p *PersistentCollection) SetReader(readkey, filepath string) error {
	var err error
	p.PersistentReader, err = NewPersistentReader(filepath, readkey, p.MaxBufferSize)
	return err
}

func (p *PersistentCollection) SetWriter(jsonkey string, outputfile *os.File) error {
	var err error
	p.PersistentWriter, err = NewPersistentWriter(jsonkey, outputfile, p.MaxBufferSize)
	return err
}

//getBufferSize returns the current maximum buffer size
//
//By default, the value is set to 50,000
func (p *PersistentCollection) GetBufferSize() int {
	return p.MaxBufferSize
}

func (p *PersistentCollection) GetReaderError() error {
	return p.PersistentReader.ErrorsQueue.Get()
}

func (p *PersistentCollection) GetWriterError() error {
	return p.PersistentWriter.ErrorsQueue.Get()
}

func (p *PersistentCollection) GetReaderKey() string {
	return p.PersistentReader.JsonKey
}

func (p *PersistentCollection) GetWriterKey() string {
	return p.PersistentWriter.JsonKey
}

//setBufferSize sets the maxBufferSize to newSize
func (p *PersistentCollection) SetBufferSize(newBufferSize int) {
	p.MaxBufferSize = newBufferSize
}

func (p *PersistentCollection) Add() {

}

func (p *PersistentCollection) Next(recordOutput any) error {
	return p.PersistentReader.NextRecord(recordOutput)
}

func (p *PersistentCollection) ReaderLength() (int, error) {
	if p.PersistentReader.Empty {
		return 0, nil
	}
	if p.PersistentReader.Length == 0 {
		for item := new(any); p.Next(item) == nil; item = new(any) {
		}
		p.PersistentReader.Reset()
		if err := p.GetReaderError(); err != nil {
			return 0, err
		}
	}
	return p.PersistentReader.Length, nil
}

func (p *PersistentCollection) ResetReader() {
	p.PersistentReader.Reset()
}

func (p *PersistentCollection) CloseReader() error {
	return p.PersistentReader.Close()
}
