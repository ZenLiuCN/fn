package fn

import (
	"bytes"
	"sync"
)

// LimitDetector detect a value is off limits
type LimitDetector[T any] func(T) bool

// ValueCreator create new value
type ValueCreator[T any] func() T

// ValuePurifier purify a Value's content
type ValuePurifier[T any] func(T) T
type SimplePool[T any] struct {
	p     sync.Pool
	reset ValuePurifier[T]
	limit LimitDetector[T]
}

/*
NewPool create with ctor ,reset function and limit detector

ctor: require creator of T

reset: required function to reset the data of T

limit: optional function that detect T is off limit, which will not put back to pool
*/
func NewPool[T any](ctor ValueCreator[T], reset ValuePurifier[T], limit LimitDetector[T]) *SimplePool[T] {
	return &SimplePool[T]{p: sync.Pool{New: func() any { return ctor() }}, reset: reset, limit: limit}
}

func (p *SimplePool[T]) Get() T {
	return p.p.Get().(T)
}
func (p *SimplePool[T]) GetPooled() *Pooled[T] {
	return &Pooled[T]{
		v:    p.p.Get().(T),
		pool: p,
	}
}
func (p *SimplePool[T]) Put(t T) {
	if p.limit != nil && p.limit(t) {
		return
	}
	p.p.Put(p.reset(t))
}

type Pooled[T any] struct {
	v     T
	freed bool
	pool  *SimplePool[T]
}

func (t *Pooled[T]) IsFreed() bool {
	return t.freed
}
func (t *Pooled[T]) Underling() T {
	return t.v
}
func (t *Pooled[T]) Free() {
	if t.freed {
		return
	}
	t.pool.Put(t.v)
	t.freed = true
	var x T
	t.v = x
}
func NewByteBufferPool() *SimplePool[*bytes.Buffer] {
	return NewPool(func() *bytes.Buffer {
		return new(bytes.Buffer)
	}, func(buffer *bytes.Buffer) *bytes.Buffer {
		buffer.Reset()
		return buffer
	}, nil)
}
func NewByteBufferPoolLimited(limit uint) *SimplePool[*bytes.Buffer] {
	i := int(limit)
	return NewPool(func() *bytes.Buffer {
		return new(bytes.Buffer)
	}, func(buffer *bytes.Buffer) *bytes.Buffer {
		buffer.Reset()
		return buffer
	}, func(buf *bytes.Buffer) bool {
		return buf.Cap() >= i
	})
}
func NewByteBufferPoolInitialized(size uint) *SimplePool[*bytes.Buffer] {
	return NewPool(func() *bytes.Buffer {
		return bytes.NewBuffer(make([]byte, 0, size))
	}, func(buffer *bytes.Buffer) *bytes.Buffer {
		buffer.Reset()
		return buffer
	}, nil)
}
func NewByteBufferPoolInitializedLimited(size uint, limit uint) *SimplePool[*bytes.Buffer] {
	i := int(limit)
	return NewPool(func() *bytes.Buffer {
		return bytes.NewBuffer(make([]byte, 0, size))
	}, func(buffer *bytes.Buffer) *bytes.Buffer {
		buffer.Reset()
		return buffer
	}, func(buf *bytes.Buffer) bool {
		return buf.Cap() >= i
	})
}
