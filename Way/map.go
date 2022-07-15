package main

import (
	"fmt"
	"sort"
)

func roChan(ch chan<- string) {
	close(ch)
}
func _close() {
	ch := make(chan string)

	close(ch)
	fmt.Println(<-ch)
	roChan(ch)
}

func mf1() {
	var a int = 1
	var b int = 2
	m := make(map[*int]byte)
	m[&a] = 'a'
	m[&b] = 'b'

	for k, v := range m {
		fmt.Println(*k, v)
	}
}

func capicity() {
	m := make(map[string]string, 2)
	m["a"] = "a"
	m["b"] = "b"
	fmt.Println(len(m))
	m["cgo"] = "cgo"
	fmt.Println(len(m))
	delete(m, "d")
	fmt.Println(len(m))
	delete(m, "cgo")
	fmt.Println(len(m))
}

type sortMap struct {
	key   int
	value int
}

func trav() {
	m := make(map[int]int)
	m[1] = 1
	m[2] = 2

	m[3] = 3
	m[4] = 4

	sm := make([]sortMap, len(m))
	i := 0
	for key := range m {
		sm[i].key = key
		i++
	}
	sort.Slice(sm, func(i, j int) bool {
		return sm[i].key < sm[j].key
	})
	for i, v := range sm {
		sm[i].value = m[v.key]
	}
	fmt.Println(sm)
}

func main() {
	//mf1()
	//capicity()
	//trav()
	_close()

}
