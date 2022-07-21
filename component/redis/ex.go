package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var client *redis.Client
func init() {
	client = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
	})
}
func ClientExample(){
	b, err := client.SetNX(ctx, "lock", "11111", 0).Result()
	if err != nil{
		panic(err)
	}else{
		fmt.Println("b, err:", b, err)
	}

	val, err := client.Get(ctx, "lock").Result()
	 if err != nil{
		panic(err)
	}else{
		fmt.Println("lock is", val)
	}
	val2, err := client.Get(ctx, "lock2").Result()
	if err == redis.Nil{
		fmt.Println("lock2 not exists")
	}else if err != nil{
		fmt.Println(err)
	}else {
		fmt.Println(val2)
	}
}

func ZsetExample(){
	b := []redis.Z{
		{
			1, "long-er",
		},
		{
			2, "xiang-er",
		},
		{
			3, "rong-er",
		},
	}
	client.ZAdd(ctx,"myset", &b[0], &b[1], &b[2])

	luaScript := redis.NewScript(`
		local element = redis.call("ZRANGE", KEYS[1], 0, 0)
		redis.call("ZREM", KEYS[1], element[1])
		return element[1]
`)
	res, err := luaScript.Run(ctx, client, []string{"myset"}, 1).Result()
	if err != nil{
		panic(err)
	}
	fmt.Println("res:", res)
}
func main1() {
	ClientExample()
	//ZsetExample()
}

