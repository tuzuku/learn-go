package main

import (
	"fmt"
	"time"

	redisutil "redisutils/redisutils"
)

func main() {
	client := redisutil.NewClient()
	err := client.Set("key", "TUZUKU", 10*time.Minute)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	val, err := client.Get("key")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Value:", val)

}
