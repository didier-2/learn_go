package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Runner struct {
	Name string
}

func (r Runner) Run(startPiontWg, wg *sync.WaitGroup) {
	defer wg.Done()
	startPiontWg.Wait()

	start := time.Now()
	fmt.Println(r.Name, "开始跑@", start)
	rand.Seed(time.Now().UnixNano())
	rand.Int()
	time.Sleep(time.Duration(rand.Uint64()%10) * time.Second)
	finish := time.Now()
	fmt.Println(r.Name, "跑到终点，用时", finish.Sub(start))
}

func main() {
	runnerCount := 10
	runners := []Runner{}

	wg := sync.WaitGroup{}
	wg.Add(runnerCount)

	startPiontWg := sync.WaitGroup{}
	startPiontWg.Add(1)

	for i := 0; i < runnerCount; i++ {
		runners = append(runners, Runner{
			Name: fmt.Sprintf("%d", i),
		})
	}
	for _, runnersItem := range runners {
		go runnersItem.Run(&startPiontWg, &wg)

	}

	fmt.Println("各就位")
	time.Sleep(1 * time.Second)
	fmt.Println("预备，跑！")

	startPiontWg.Done()
	wg.Wait()
	fmt.Println("赛跑结束")
}
