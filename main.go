package main

import (
	"fmt"
	"inMemoryCash/inMemoryCash"
)

func main() {

	cache := inMemoryCash.NewCache()

	cache.Set("key1", "Value1")
	fmt.Println(cache.Get("key1"))

	cache.Delete("key1")
	fmt.Println(cache.Get("key1"))

}
