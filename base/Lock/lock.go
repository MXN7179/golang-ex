package main

import (
	"fmt"
	"sync"
)

var m = sync.Mutex{}

func f1() {
	m.Lock()
	defer m.Unlock()
	fmt.Println("func f1")
	f2()
}

func f2() {
	m.Lock()
	defer m.Unlock()
	fmt.Println("func f2")
}

func main() {
	f1()
}
