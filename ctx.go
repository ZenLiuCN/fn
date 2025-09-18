package fn

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func ContextSignal() (ctx context.Context, cc context.CancelFunc, ch chan os.Signal) {
	ctx, cc = context.WithCancel(context.Background())
	ch = make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	return
}
func ContextSignalGo() (context.Context, context.CancelFunc) {
	ctx, cc, ch := ContextSignal()
	go func() {
		<-ch
		cc()
	}()
	return ctx, cc
}
