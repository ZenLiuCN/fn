package fn

import (
	"context"
	"os"
	"os/signal"
)

// WithSignal run a contextual function with notify by signals, the returning context.CancelFunc is optional to invoke.
func WithSignal(fn func(ctx context.Context), signals ...os.Signal) context.CancelFunc {
	ctx, cc := signal.NotifyContext(context.Background(), signals...)
	go func() {
		fn(ctx)
		cc()
	}()
	return cc
}
