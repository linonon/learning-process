package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

func main() {
	c := NewClient()
	test(c)
}

func NewClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", //no password
		DB:       0,  // use defalt db
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client
}

func test(c *redis.Client) {
	err := c.Set("linonon", "1", 0).Err() // 0 為 永不過時
	if err != nil {
		panic(err)
	}
	c.LPush("linononLpush", "Lpushing")

	val, err := c.Get("linonon").Result() // Get操作，並獲得返回值
	if err != nil {
		panic(err)
	}
	fmt.Println("linonon: ", val)

	val2, err := c.Get("key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}
