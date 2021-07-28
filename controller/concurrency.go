package controller

import (
	"fmt"
	"sync"
	"time"
)

/*
Simple Demo Concurrency
*/
func GetConcurrency() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {
		defer waitGroup.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Hello")
	}()
	go func() {
		defer waitGroup.Done()
		fmt.Println("First")
	}()
	fmt.Println("Last")
	waitGroup.Wait()
}

/*
Demo for channel
*/
func ChannelDemo() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)
	wg.Add(1)
	// <- means the direction of data flow in the channel; without it its bidirectional
	go func(ch <-chan int, wg *sync.WaitGroup) {
		if msg, ok := <-ch; ok {
			fmt.Println(msg, ok)
		}
		if msg, ok := <-ch; ok {
			fmt.Println(msg, ok)
		}
		for msg := range ch {
			fmt.Println(msg)
		}
		// Channel closed
		if msg, ok := <-ch; ok {
			fmt.Println(msg, ok)
		}
		wg.Done()
	}(ch, wg)
	wg.Add(1)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 42
		ch <- 27
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch) // Cant reopen a closed channel; Cant close channel is receiving side
		// If close is missing then for msg :=range will cause panic
		wg.Done()
	}(ch, wg)
	wg.Wait()
}
