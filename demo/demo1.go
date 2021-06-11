package main

func main() {
	concurrency1()
}

// 问题多多
func concurrency1() {
	for i := 0; i < 5; i++ {
		go func() {
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

	printResult()
}
