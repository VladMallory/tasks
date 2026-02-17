package counter

import "sync/atomic"

type AtomicCounter struct {
	value atomic.Int64
}

func NewAtomicCounter() *AtomicCounter {
	return &AtomicCounter{}
}

// Increment атомарно увеличивает значение счетчика.
func (a *AtomicCounter) Increment() {
	a.value.Add(1)
}

// Value атомарно возвращает текущее значение счетчика.
func (a *AtomicCounter) Value() int64 {
	return a.value.Load()
}
