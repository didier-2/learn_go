package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	startTime := time.Now()
	maxNum := 2000000
	result := make(chan int, maxNum)
	wg := sync.WaitGroup{}
	workerNumber := 16
	wg.Add(workerNumber)
	baseNumberCh := make(chan int, 1024)
	for i := 0; i < workerNumber; i++ {
		go func() {
			defer wg.Done()
			for oNumber := range baseNumberCh {
				if isPrime(oNumber) {
					result <- oNumber
				}
			}
		}()
	}
	for num := 2; num <= maxNum; num++ {
		baseNumberCh <- num
	}
	close(baseNumberCh)
	wg.Wait()

	finishTime := time.Now()
	fmt.Println(len(result))
	fmt.Println("共耗时：", finishTime.Sub(startTime))
}
func isPrime(num int) (isPrime bool) {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			isPrime = false
			return
		}
	}
	isPrime = true
	return
}
