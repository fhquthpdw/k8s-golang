package main

import (
	"context"
	"sync"
	"sync/atomic"

	"golang.org/x/sync/semaphore"
)

type returnResult struct {
	url    string
	result string
	err    error
}

func main() {
	var ch = make(chan returnResult)

	mwg := &sync.WaitGroup{}
	mwg.Add(2)
	go concurrency7(ch, mwg)
	go getResult(ch, mwg)
	mwg.Wait()

	printResult()
}

func getResult(ch chan returnResult, wg7 *sync.WaitGroup) {
	defer wg7.Done()

	for v := range ch {
		if v.err == nil {
			atomic.AddInt32(&success, 1)
			successResult[v.url] = v.result
		} else {
			atomic.AddInt32(&failed, 1)
			failedError[v.url] = v.err.Error()
		}
	}
}

func concurrency7(ch chan returnResult, wg7 *sync.WaitGroup) {
	defer wg7.Done()

	var goroutineCnt int64 = 10
	ctx := context.Background()
	sema := semaphore.NewWeighted(goroutineCnt)

	for i := 0; i < 300; i++ {
		sema.Acquire(ctx, 1)

		go func(i int) {
			defer sema.Release(1)

			str, err := sendRequest(urlArr[i])
			ch <- returnResult{
				url:    urlArr[i],
				result: str,
				err:    err,
			}
		}(i)
	}

	sema.Acquire(ctx, goroutineCnt)
	close(ch) // if close then deadlock will occurred
}
