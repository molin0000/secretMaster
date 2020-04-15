package ui

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestUI(t *testing.T) {
	StartUI()
	time.Sleep(10 * time.Second)
	StopUI()
	time.Sleep(10 * time.Second)
}

func TestContext(t *testing.T) {
	ctx, stop := context.WithCancel(context.Background())
	stop()

	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("exit")
				return
			default:
				fmt.Println("sleep")
				time.Sleep(time.Second * 1)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("stop")
	time.Sleep(time.Second * 3)
}
