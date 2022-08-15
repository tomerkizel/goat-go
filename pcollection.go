package pcollection

type PersistentCollection struct {
	MaxBufferSize    int               `json:"maxBufferSize"`
	PersistentReader *PersistentReader `json:"persistentReader"`
	PersistentWriter *PersistentWriter `json:"persistentWriter"`
}

func NewPersistentCollection(readfilepath, writefilepath, readkey string) (*PersistentCollection, error) {
	self := NewEmptyPersistentCollection()
	err := self.SetReader(readfilepath, readkey)
	if err != nil {
		return nil, err
	}
	err = self.SetWriter(writefilepath)
	if err != nil {
		return nil, err
	}
	return self, nil
}

func NewEmptyPersistentCollection() *PersistentCollection {
	self := PersistentCollection{}
	self.SetAndGetNewBufferSize(50000)
	return &self
}

func (p *PersistentCollection) SetReader(filepath, readkey string) error {
	var err error
	p.PersistentReader, err = NewPersistentReader(filepath, readkey, p.MaxBufferSize)
	return err
}

func (p *PersistentCollection) SetWriter(filepath string) error {
	var err error
	p.PersistentWriter, err = NewPersistentWriter(filepath, p.MaxBufferSize)
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

//setBufferSize sets the maxBufferSize to newSize
func (p *PersistentCollection) SetAndGetNewBufferSize(newBufferSize int) {
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
