package main

import (
	"fmt"
	"sync"
	"time"
)

func Producer(wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		c <- i
		fmt.Printf("Produce %d\n", i)
		time.Sleep(time.Second * 1)
	}
	close(c)
}

func Consumer(wg *sync.WaitGroup, c chan int) {
	defer wg.Done()
	for i := range c {
		fmt.Printf("Consume %d\n", i)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	c := make(chan int, 10)
	var wait sync.WaitGroup
	wait.Add(2)
	go Producer(&wait, c)
	go Consumer(&wait, c)
	wait.Wait()
}
