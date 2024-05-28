package fsm

import (
	"context"
	"os"
)

type StateFn[T any, S ~func(ctx context.Context, args T) (T, S, error)] func(ctx context.Context, args T) (T, S, error)

func EvolveFn[T any, S ~func(ctx context.Context, args T) (T, S, error)](ctx context.Context, args T, current S) (o T, err error) {
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

type Decider[K comparable, T any, S ~func(ctx context.Context, args T) (T, S, error)] map[K]S
type Decision[K comparable, T any, S ~func(ctx context.Context, args T) (T, S, error)] func(ctx context.Context, args T) (T, S, error)

// EvolveDecision process strategies from start with args. returns [os.ErrNotExist] if the start state not exists
func EvolveDecision[K comparable, T any, S ~func(ctx context.Context, args T) (T, S, error), D ~map[K]S](s D, ctx context.Context, start K, args T) (o T, err error) {
	current, ok := s[start]
	if !ok {
		return o, os.ErrNotExist
	}
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
