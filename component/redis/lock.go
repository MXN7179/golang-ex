package main

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"math/rand"
	"strconv"
	"time"
)

const (
	unlockStr = `
	if redis.call("get", KEYS[1]) == ARGV[1] then
		return redis.call("del", KEYS[1])
	else
		return 0
	end
`
)

var(
 	//ctx = context.Background()
	//client *redis.Client

	value int
	key string
)
func init(){
	//initRedis()
	initA()
}
func initRedis() {
	client = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
	})
}
func initA() {
	rand.Seed(time.Now().Unix())
	value = rand.Intn(1000)
	key = "key1"
}

func main() {
	Lock()
	defer UnLock()
	// do ...
	time.Sleep(time.Second * 3)

}

func Lock(){
	succ, err := client.SetNX(ctx, key, strconv.Itoa(value), time.Second * 5).Result()
	if err != nil{
		panic(err)
	}else if succ == false{
		fmt.Printf("key has exists\n")
	}else{
		fmt.Printf("key set success\n")
	}
}

func UnLock(){
	unlockScript := redis.NewScript(unlockStr)
	res, err := unlockScript.Run(ctx, client, []string{key}, strconv.Itoa(value)).Result()
	fmt.Printf("res: %v err: %v\n", res, err)
}