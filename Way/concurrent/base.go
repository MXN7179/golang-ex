package main

import (
	"fmt"
	"math/rand"
	"time"
)

func countGoroutine() {
	i := 0
	for {
		i++
		go func(i int) {
			fmt.Println(i)
			for {

			}
		}(i)
	}
}

const MaxValue = 100
const MaxOutStanding = 2

type Request struct {
	v int
}

var ch = make(chan *Request, MaxValue)

func push(ch chan *Request) {
	for i := 0; ; i++ {
		ch <- &Request{v: i}
	}
}

func handle(f int, reqs chan *Request) {
	for r := range reqs {
		fmt.Println(f, r.v)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(50)))
	}
}
func Serve(clientRequests chan *Request, quit chan bool) {
	for i := 0; i < MaxOutStanding; i++ {
		go handle(i, clientRequests)
	}
	<-quit
}
func Do() {
	// send to chan
	go push(ch)

	quit := make(chan bool)
	Serve(ch, quit)
}
func main() {
	rand.Seed(time.Now().Unix())
	Do()
}
