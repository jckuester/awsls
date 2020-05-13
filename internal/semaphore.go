package internal

// Semaphore is a thin wrapper around a channel for using it as a semaphore.
type Semaphore chan struct{}

// NewSemaphore creates a semaphore that allows up to a given limit of
// simultaneous acquisitions.
func NewSemaphore(n int) Semaphore {
	if n == 0 {
		panic("semaphore with limit 0")
	}

	ch := make(chan struct{}, n)
	return ch
}

// Acquire is used to acquire an available slot. Blocks until available.
func (s Semaphore) Acquire() {
	s <- struct{}{}
}

// Release is used to return a slot. Acquire must be called as a pre-condition.
func (s Semaphore) Release() {
	select {
	case <-s:
	default:
		panic("release without an acquire")
	}
}
