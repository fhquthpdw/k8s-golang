package main

import "sync/atomic"

func main() {
	concurrency4()
}

// 解决i++原子操作问题
func concurrency4() {
	for i := 0; i < 300; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done() // same as defer wg.Add(-1)

			url := urlArr[i]
			if str, err := sendRequest(url); err == nil {
				atomic.AddInt32(&success, 1)
				successResultSync.Store(url, str)
			} else {
				atomic.AddInt32(&failed, 1)
				failedErrorSync.Store(url, err.Error())
			}
		}()
	}
	wg.Wait()

	printResult3()
}
