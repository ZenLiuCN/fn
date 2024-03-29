package fn

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

// WithSignal run a contextual function with notify by signals, the returning context.CancelFunc is optional to invoke.
func WithSignal(fn func(ctx context.Context), signals ...os.Signal) context.CancelFunc {
	if len(signals) == 0 {
		signals = []os.Signal{
			syscall.SIGINT,
			syscall.SIGTERM,
		}
	}
	ctx, cc := signal.NotifyContext(context.Background(), signals...)
	go func() {
		fn(ctx)
		cc()
	}()
	return cc
}

// WithSignalA use signal as context to run a function
// @return context.CancelFunc to cancel execution
// @return <-chan A to receive result, it will close after execution
func WithSignalA[A any](fn func(ctx context.Context) A, signals ...os.Signal) (context.CancelFunc, <-chan A) {
	if len(signals) == 0 {
		signals = []os.Signal{
			syscall.SIGINT,
			syscall.SIGTERM,
		}
	}
	ctx, cc := signal.NotifyContext(context.Background(), signals...)
	ch := make(chan A, 1)
	go func() {
		defer close(ch)
		ch <- fn(ctx)
		cc()
	}()
	return cc, ch
}
