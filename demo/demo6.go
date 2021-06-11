package main

import (
	"context"
	"sync/atomic"

	"golang.org/x/sync/semaphore"
)

func main() {
	concurrency6()
}

func concurrency6() {
	var goroutineCnt int64 = 10
	ctx := context.Background()
	sema := semaphore.NewWeighted(goroutineCnt)

	for i := 0; i < 300; i++ {
		sema.Acquire(ctx, 1)

		go func(i int) {
			defer sema.Release(1)

			url := urlArr[i]
			if str, err := sendRequest(url); err == nil {
				atomic.AddInt32(&success, 1)
				successResultSync.Store(url, str)
			} else {
				atomic.AddInt32(&failed, 1)
				failedErrorSync.Store(url, err.Error())
			}
		}(i)
	}
	sema.Acquire(ctx, goroutineCnt)

	printResult3()
}
