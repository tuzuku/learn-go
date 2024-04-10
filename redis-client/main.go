package main

import ( 
	"fmt"
	"github.com/go-redis/redis"
	"context"
)


func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		Password: "",
		DB: 0,
	})

	pong,err := rdb.Ping().Result()
	if err != nil {
		fmt.Println("Error:",err)
		return
	}
	fmt.Println("PING response: ",pong)

	err = rdb.Set( "key", "value", 0).Err()
	if err != nil {
		fmt.Println("Error:",err)
		return
	}

	val,err := rdb.Get( "key").Result()
	if err != nil {
		fmt.Println("Error:",err)
		return
	}
	fmt.Println("key",val)

	_,err = rdb.Del( "key").Result()
	if err != nil {
		fmt.Println("Error:",err)
		return
	}
	fmt.Println("key deleted")


}