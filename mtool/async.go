package mtool

import (
	"context"
	"fmt"
	"log"
)

func FunCall(ctx context.Context, fn func()) {
	nextChan := make(chan struct{})
	go func(c chan<- struct{}, f func()) {
		defer func() {
			if p := recover(); p != nil {
				log.Fatalf(fmt.Sprintf("AsyncFunCall recover, p: %v", fmt.Sprint(p)))
			}
		}()
		f()
		close(c)
	}(nextChan, fn)
	
	select {
	case <-ctx.Done():
		return
	case <-nextChan:
		return
	}
}
