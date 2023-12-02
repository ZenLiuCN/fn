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
