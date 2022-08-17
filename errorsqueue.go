package goat

const MIN_SIZE int = 1

type ErrorsQueue struct {
	errors chan error
	empty  bool
}

func NewErrorsQueue(queuesize int) *ErrorsQueue {
	self := ErrorsQueue{}
	if queuesize < MIN_SIZE {
		queuesize = MIN_SIZE
	}
	self.errors = make(chan error, queuesize)
	self.empty = true
	return &self
}

func (eq *ErrorsQueue) Add(err error) {
	if err != nil {
		eq.errors <- err
		eq.empty = false
	}
}

func (eq *ErrorsQueue) Get() error {
	if eq.empty {
		return nil
	}
	err := <-eq.errors
	if len(eq.errors) == 0 {
		eq.empty = true
	}
	return err
}
