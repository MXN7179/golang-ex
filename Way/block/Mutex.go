package main

import (
	"fmt"
	"sync"
)

var mu = sync.Mutex{}

func main() {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("main-")
	A()

}

func A() {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("A")
}
