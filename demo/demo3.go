package main

func main() {
	concurrency3()
}

// 解决map写冲突问题
func concurrency3() {
	for i := 0; i < 300; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done() // same as defer wg.Add(-1)

			url := urlArr[i]
			if str, err := sendRequest(url); err == nil {
				success++
				successResultSync.Store(url, str)
			} else {
				failed++
				failedErrorSync.Store(url, err.Error())
			}
		}()
	}
	wg.Wait()

	printResult3()
}
