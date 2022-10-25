package main

import (
	"fmt"
	"github.com/Ixander/ninjaGo/week1/cache1"
)

func main() {

	cache := cache1.New()

	cache.Set("userId", 42)

	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId = cache.Get("userId")

	fmt.Println(userId)
}
