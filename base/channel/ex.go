package main

import (
	"fmt"
	"math/rand"
	"time"
)

// N senderr, 1 reciver
func Nto1() {
	rand.Seed(time.Now().Unix())

	dataCh := make(chan int, 10)
	stopCh := make(chan int, 1)

	const Max int = 100
	const MaxSender int = 100

	for i := 0; i < MaxSender; i++ {
		go func() {
			for {
				select {
				case <-stopCh:
					return
				case dataCh <- rand.Intn(Max):
				}
			}
		}()
	}

	for value := range dataCh {
		fmt.Println(value)
		if value == Max-1 {
			fmt.Println("stopCh")
			close(stopCh)
			break
		}
	}
	select {
	case <-time.After(time.Hour):
	}
}

func main1() {
	Nto1()
}
