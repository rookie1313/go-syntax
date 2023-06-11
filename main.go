package main

import (
	"fmt"
	"sync"
)

func main() {
	c := make(chan int, 100)

	wg := sync.WaitGroup{}
	wg.Add(2)
	//2 producers
	go func() {
		defer wg.Done()
		for i := 0; i < 90; i++ {
			c <- i
		}
	}()
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	//for block main,so use struct{}
	block := make(chan struct{}, 0)
	//1 consumer
	go func() {
		sum := 0
		for {
			a, ok := <-c
			if ok {
				sum += a
			} else {
				break
			}
		}
		fmt.Printf("sum is %d", sum)
		//add element,the purpose is unblocking this channel
		block <- struct{}{}
	}()

	wg.Wait()
	close(c)
	//if no element,block
	<-block
}
