package main

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"
)

func createPost(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	return nil
}

var wg sync.WaitGroup

func _p2() {
	defer log.Println("p2")
	//log.Fatalf("p2 fatal")
	panic("p2 panic")
}
func _p1() {
	defer log.Println("p1")
	wg.Add(1)
	go _p2()
	log.Println("p1 end")
	wg.Wait()
}

func startAt() {
	startAt := time.Now()
	//defer fmt.Println(time.Since(startAt)) // 直接拷贝startAt 传入defer
	defer func() { fmt.Println(time.Since(startAt)) }() // 拷贝给defer的是 *func

	time.Sleep(time.Second)
}

func main() {
	//_p1()
	startAt()
}
