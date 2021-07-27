package controller

import (
	"fmt"
	"sync"
	"time"
)

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
