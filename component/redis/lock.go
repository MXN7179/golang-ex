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
	renewalStr = `
	  if redis.call("get", KEYS[1]) == ARGV[1] then
		return redis.call("expire", ARGV[2])
	  else
		return 0
	  end
`
)

var(
	value int
	key string

	renewalCh = make(chan struct{})
)
func init(){
	rand.Seed(time.Now().Unix())
	value = rand.Intn(1000)
	key = "key1"
}
func WatchDog(requestId string){
	ticker := time.NewTicker(3 * time.Second)
	renewalScript := redis.NewScript(renewalStr)

	for {
		select {
		case <- ticker.C:
			fmt.Println("续期5s")
			renewalScript.Run(ctx, client, []string{key}, requestId, 5)  // 续期5秒
		case <- renewalCh:
			return
		}
	}
}
func Lock(){
	succ, err := client.SetNX(ctx, key, strconv.Itoa(value), time.Second * 5).Result()
	if err != nil{
		panic(err)
	}else if succ == false{
		fmt.Printf("key has exists\n")
	}else{
		fmt.Printf("key set success\n")
		go WatchDog(strconv.Itoa(value))	 // 开启看门狗
	}
}

func UnLock(){
	unlockScript := redis.NewScript(unlockStr)
	res, err := unlockScript.Run(ctx, client, []string{key}, strconv.Itoa(value)).Result()
	fmt.Printf("res: %v err: %v\n", res, err)
	if err != nil{
		fmt.Println("unlock fail")
	}else if res == 0 {
		fmt.Println("lock not exists")
	}else{
		fmt.Println("renewal <- ")
		renewalCh <- struct{}{}
	}
}

func main() {
	Lock()
	defer UnLock()
	// do ...
	time.Sleep(time.Second * 8)


}
