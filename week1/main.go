package main

import (
	"fmt"
	"github.com/Ixander/ninjaGo/week1/cache1"
	"log"
	"time"
)

func main() {

	cache := cache1.New()

	//cache.Set("userId", 42)
	//
	//userId := cache.Get("userId")
	//
	//fmt.Println(userId)
	//
	//cache.Delete("userId")
	//userId = cache.Get("userId")
	//
	//fmt.Println(userId)

	cache.Set("userId", 42, time.Second*2)
	userId, err := cache.Get("userId")
	if err != nil { // err == nil
		log.Fatal(err)
	}
	fmt.Println(userId) // Output: 42

	time.Sleep(time.Second * 3) // прошло 5 секунд

	userId, err = cache.Get("userId")
	if err != nil { // err != nil
		log.Fatal(err) // сработает этот код
	}
	//fmt.Println(userId)
}
