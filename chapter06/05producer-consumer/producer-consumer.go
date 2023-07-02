package main

import (
	"fmt"
	"sync"
	"time"
)

type Store struct {
	DateCount int
	Max       int
	lock      sync.Mutex
}

type Producer struct {
}

func (Producer) Produce(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.DateCount == s.Max {
		fmt.Println("生产者看到库存满了，不生产")
		return
	}
	fmt.Println("开始生产+1")
	s.DateCount++
}

type Consumer struct {
}

func (Consumer) Consume(s *Store) {
	s.lock.Lock()
	defer s.lock.Unlock()
	if s.DateCount == 0 {
		fmt.Println("消费者看到没库存了，不消费")
		return
	}
	fmt.Println("消费者消费-1")
	s.DateCount--
}

func main() {
	s := &Store{
		Max: 10,
	}
	pCount, cCount := 50, 50
	for i := 0; i < pCount; i++ {
		go func() {
			for {
				time.Sleep(100 * time.Millisecond)
				Producer{}.Produce(s)
			}
		}()
	}
	for i := 0; i < cCount; i++ {
		go func() {
			for {
				time.Sleep(100 * time.Millisecond)
				Consumer{}.Consume(s)
			}
		}()
	}
	time.Sleep(1 * time.Second)
	fmt.Println(s.DateCount)
}
