package main

import "fmt"

func _s() {
	c := make(chan int)
	//var i int

	go func() {
		for {
			fmt.Println(<-c)
		}
	}()

	for {
		select {
		case c <- 0:
		case c <- 1:
		}
	}

}

func main() {
	_s()

	//select {}
}
