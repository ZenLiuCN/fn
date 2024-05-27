package fn

import "context"

type State[T any] func(ctx context.Context, args T) (T, State[T], error)

func Evolve[T any](ctx context.Context, args T, current State[T]) (_ T, err error) {
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

type FSM[K comparable, T any] map[K]Evolves[K, T]
type Evolves[K comparable, T any] func(fsm FSM[K, T], ctx context.Context, args T) (T, Evolves[K, T], error)

func (s FSM[K, T]) Evolve(ctx context.Context, start K, args T) (_ T, err error) {
	current := s[start]
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
