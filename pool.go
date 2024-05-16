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
type Pool[T any] struct {
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
func NewPool[T any](ctor ValueCreator[T], reset ValuePurifier[T], limit LimitDetector[T]) *Pool[T] {
	return &Pool[T]{p: sync.Pool{New: func() any { return ctor() }}, reset: reset, limit: limit}
}

func (p *Pool[T]) Get() T {
	return p.p.Get().(T)
}
func (p *Pool[T]) GetPooled() *Pooled[T] {
	return &Pooled[T]{
		v:    p.p.Get().(T),
		pool: p,
	}
}
func (p *Pool[T]) Put(t T) {
	if p.limit != nil && p.limit(t) {
		return
	}
	p.p.Put(p.reset(t))
}

type Pooled[T any] struct {
	v     T
	freed bool
	pool  *Pool[T]
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
func NewByteBufferPool() *Pool[*bytes.Buffer] {
	return NewPool(func() *bytes.Buffer {
		return new(bytes.Buffer)
	}, func(buffer *bytes.Buffer) *bytes.Buffer {
		buffer.Reset()
		return buffer
	}, nil)
}
func NewByteBufferPoolLimited(limit uint) *Pool[*bytes.Buffer] {
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
func NewByteBufferPoolInitialized(size uint) *Pool[*bytes.Buffer] {
	return NewPool(func() *bytes.Buffer {
		return bytes.NewBuffer(make([]byte, 0, size))
	}, func(buffer *bytes.Buffer) *bytes.Buffer {
		buffer.Reset()
		return buffer
	}, nil)
}
func NewByteBufferPoolInitializedLimited(size uint, limit uint) *Pool[*bytes.Buffer] {
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

type TinyPool[T any] sync.Pool

func NewTinyPool[T any](newFunc func() T) *TinyPool[T] {
	return &TinyPool[T]{New: func() any { return newFunc() }}
}
func (t *TinyPool[T]) Put(v T) {
	((*sync.Pool)(t)).Put(v)
}
func (t *TinyPool[T]) Get() (v T) {
	return ((*sync.Pool)(t)).Get().(T)
}
