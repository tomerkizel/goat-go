package goat

type PersistentCollection struct {
	MaxBufferSize    int `json:"maxBufferSize"`
	persistentReader *PersistentReader
	persistentWriter *PersistentWriter
}

func NewPersistentCollection(readfilepath string, writefullfile bool, readkey, writekey string) (*PersistentCollection, error) {
	self := newEmptyPersistentCollection()
	err := self.setReader(readkey, readfilepath)
	if err != nil {
		return nil, err
	}
	err = self.setWriter(writekey, writefullfile)
	if err != nil {
		return nil, err
	}
	return self, nil
}

func newEmptyPersistentCollection() *PersistentCollection {
	self := PersistentCollection{}
	self.MaxBufferSize = 50000
	return &self
}

func (p *PersistentCollection) setReader(readkey, filepath string) error {
	var err error
	p.persistentReader, err = newPersistentReader(filepath, readkey, p.MaxBufferSize)
	return err
}

func (p *PersistentCollection) setWriter(jsonkey string, iscompletefile bool) error {
	var err error
	p.persistentWriter, err = newPersistentWriter(jsonkey, iscompletefile, p.MaxBufferSize)
	return err
}

//getBufferSize returns the current maximum buffer size
//
//By default, the value is set to 50,000
func (p *PersistentCollection) GetBufferSize() int {
	return p.MaxBufferSize
}

//setBufferSize sets the maxBufferSize to newSize
func (p *PersistentCollection) SetBufferSize(newBufferSize int) {
	p.MaxBufferSize = newBufferSize
}

func (p *PersistentCollection) GetReaderError() error {
	return p.persistentReader.getError()
}

func (p *PersistentCollection) GetWriterError() error {
	return p.persistentWriter.getError()
}

func (p *PersistentCollection) GetReaderKey() string {
	return p.persistentReader.GetJsonKey()
}

func (p *PersistentCollection) GetWriterKey() string {
	return p.persistentWriter.GetJsonKey()
}

func (p *PersistentCollection) Add(recordOutput any) {
	p.persistentWriter.write(recordOutput)
}

func (p *PersistentCollection) Next(recordOutput any) error {
	return p.persistentReader.nextRecord(recordOutput)
}

func (p *PersistentCollection) ReaderLength() (int, error) {
	if p.persistentReader.IsEmpty() {
		return 0, nil
	}
	if p.persistentReader.length == 0 {
		for item := new(any); p.Next(item) == nil; item = new(any) {
		}
		p.persistentReader.reset()
		if err := p.GetReaderError(); err != nil {
			return 0, err
		}
	}
	return p.persistentReader.length, nil
}

func (p *PersistentCollection) ResetReader() {
	p.persistentReader.reset()
}

func closeReader(p *PersistentCollection) error {
	return p.persistentReader.close()
}

func closeWriter(p *PersistentCollection) error {
	return p.persistentWriter.close()
}

func (p *PersistentCollection) Close() error {
	err := closeReader(p)
	if err == nil {
		return err
	}
	err = closeWriter(p)
	return err
}
