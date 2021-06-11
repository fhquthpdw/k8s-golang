package main

import (
	"errors"
	"fmt"
	"hash/crc32"
	"sync"
)

const urlTotal = 10000

var urlArr = generateUrls()
var success int32
var failed int32
var successResult = make(map[string]string)
var failedError = make(map[string]string)
var wg = sync.WaitGroup{}

var successResultSync = sync.Map{}
var failedErrorSync = sync.Map{}

func sendRequest(url string) (str string, err error) {
	crc32q := crc32.MakeTable(0xD5828281)
	i := crc32.Checksum([]byte(url), crc32q)
	if i%3 == 0 {
		str = url
		err = nil
	} else {
		str = ""
		err = errors.New("failed: " + url)
	}
	return
}

func generateUrls() (r []string) {
	for i := 0; i < urlTotal; i++ {
		data := []byte(string(i))
		url := string(data)
		r = append(r, url)
	}
	return
}

func getSyncMapLen(m sync.Map) (i int) {
	m.Range(func(_, _ interface{}) bool {
		i++
		return true
	})
	return
}

func printResult() {
	fmt.Println(success)
	fmt.Println(failed)
	fmt.Println(success + failed)

	fmt.Println()

	fmt.Println(len(successResult))
	fmt.Println(len(failedError))
	fmt.Println(len(successResult) + len(failedError))
}

func printResult3() {
	fmt.Println(success)
	fmt.Println(failed)
	fmt.Println(success + failed)

	fmt.Println()

	fmt.Println(getSyncMapLen(successResultSync))
	fmt.Println(getSyncMapLen(failedErrorSync))
	fmt.Println(getSyncMapLen(successResultSync) + getSyncMapLen(failedErrorSync))
}
