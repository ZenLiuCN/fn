package fsm

import (
	"context"
	"os"
)

type StateFn[T any] func(ctx context.Context, args T) (T, StateFn[T], error)

func Evolve[T any](ctx context.Context, args T, current StateFn[T]) (_ T, err error) {
	for {
		if ctx.Err() != nil {
			return args, ctx.Err()
		}
		args, current, err = current(ctx, args)
		if err != nil {
			return args, err
		}
		if current == nil {
			return args, nil
		}
	}
}

type Strategies[K comparable, T any] map[K]Strategy[K, T]
type Strategy[K comparable, T any] func(fsm Strategies[K, T], ctx context.Context, args T) (T, Strategy[K, T], error)

// Evolve process strategies from start with args. returns [os.ErrNotExist] if the start state not exists
func (s Strategies[K, T]) Evolve(ctx context.Context, start K, args T) (o T, err error) {
	current, ok := s[start]
	if !ok {
		return o, os.ErrNotExist
	}
	for {
		if ctx.Err() != nil {
			return args, ctx.Err()
		}
		args, current, err = current(s, ctx, args)
		if err != nil {
			return args, err
		}
		if current == nil {
			return args, nil
		}
	}
}
