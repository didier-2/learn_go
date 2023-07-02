package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		//countDict()
		//countDictGoroutineSafe()
		//countDictGoroutineForgetUnlock()
		//countDictLockPrice()
		countDictGoroutineSafeRW()
	}

}

func countDict() {
	totalNum := 0
	wg := sync.WaitGroup{}
	wg.Add(5000)
	for i := 0; i < 5000; i++ {
		go func() {
			defer wg.Done()
			totalNum += 100 //totalNum=totalNum+100//注意这里有重复覆盖的问题
		}()
	}
	wg.Wait()
	fmt.Println("预计有：", 100*5000)
	fmt.Println("总字数：", totalNum)
}

func countDictGoroutineSafe() {
	totalNum := 0
	totalNumlock := sync.Mutex{}

	wg := sync.WaitGroup{}
	wg.Add(5000)
	for i := 0; i < 5000; i++ {
		go func() {
			defer wg.Done()
			totalNumlock.Lock()
			totalNum += 100 //totalNum=totalNum+100//注意这里有重复覆盖的问题
			totalNumlock.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("预计有：", 100*5000)
	fmt.Println("总字数：", totalNum)
}

func countDictGoroutineSafeRW() {
	totalNum := 0
	totalNumlock := sync.RWMutex{}

	wg := sync.WaitGroup{}
	wg.Add(5000)

	go func() {
		result := map[int]struct{}{}
		for {
			totalNumlock.Lock()
			//fmt.Println(totalNum)
			result[totalNum] = struct{}{}
			totalNumlock.Unlock()
		}

		fmt.Println("result", result)
	}()
	for i := 0; i < 5000; i++ {
		go func() {
			defer wg.Done()
			totalNumlock.Lock()
			totalNum += 100 //totalNum=totalNum+100//注意这里有重复覆盖的问题
			totalNumlock.Unlock()
		}()
	}
	wg.Wait()
	time.Sleep(5 * time.Second)
	fmt.Println("预计有：", 100*5000)
	fmt.Println("总字数：", totalNum)
}
func countDictGoroutineForgetUnlock() {
	totalNum := 0
	totalNumlock := sync.Mutex{}

	wg := sync.WaitGroup{}
	wg.Add(5000)
	for i := 0; i < 5000; i++ {
		go func() {
			defer wg.Done()
			totalNumlock.Lock()
			defer totalNumlock.Unlock()

			totalNum += 100 //totalNum=totalNum+100//注意这里有重复覆盖的问题
			//totalNumlock.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println("预计有：", 100*5000)
	fmt.Println("总字数：", totalNum)
}
func countDictLockPrice() {
	totalNum := 0
	totalNumlock := sync.Mutex{}

	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(pageNum int) {
			defer wg.Done()

			totalNumlock.Lock()
			totalNum += 100 //totalNum=totalNum+100//注意这里有重复覆盖的问题

			if pageNum == 3 {
				time.Sleep(3 * time.Second)
			}
			totalNumlock.Unlock()
		}(i)
	}
	wg.Wait()
	fmt.Println("预计有：", 100*5000)
	fmt.Println("总字数：", totalNum)
}
