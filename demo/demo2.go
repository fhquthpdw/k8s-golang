package main

func main() {
	concurrency2()
}

func concurrency2() {
	// raise to 30
	for i := 0; i < 3; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done() // same as defer wg.Add(-1)

			url := urlArr[i]
			if str, err := sendRequest(url); err == nil {
				success++
				successResult[url] = str
			} else {
				failed++
				failedError[url] = err.Error()
			}
		}()
	}
	wg.Wait()

	printResult()
}
