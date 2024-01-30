package cancel

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func isCanceled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}

func isCanceledCtx(ctx context.Context) bool {
	select {
	case <-ctx.Done():
		return true
	default:
		return false
	}
}

func cancel_1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

func cancel_2(cancelChan chan struct{}) {
	close(cancelChan)
}

func TestContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	for i := 0; i < 5; i++ {
		go func(i int, ctx context.Context) {
			for {
				if isCanceledCtx(ctx) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canceled")
		}(i, ctx)
	}
	cancel()
	time.Sleep(time.Second * 1)
}
func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelChan chan struct{}) {
			for {
				if isCanceled(cancelChan) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canceled")
		}(i, cancelChan)
	}
	cancel_2(cancelChan)
	time.Sleep(time.Second * 1)
}
