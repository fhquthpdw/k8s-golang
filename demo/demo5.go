package main

import "sync/atomic"

func main() {
	concurrency5()
}

// 解决闭包参数传递问题
func concurrency5() {
	for i := 0; i < 300; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done() // same as defer wg.Add(-1)

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
	wg.Wait()

	printResult3()
}
