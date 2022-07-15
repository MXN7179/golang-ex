package main

type Request int

const MaxCount = 3

var sem = make(chan int, MaxCount)

func process(req *Request) {}
func serve(queue chan *Request) { // queue 缓冲信道
	for req := range queue {
		sem <- 1
		go func(req *Request) {
			process(req)
			<-sem
		}(req)
	}
}
