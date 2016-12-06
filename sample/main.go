package main

import (
	"cache"
	"fmt"
	"time"
)

func main() {
	defaultExpiration, _ := time.ParseDuration("0.5h")
	gcInterval, _ := time.ParseDuration("3s")
	c := cache.NewCache(defaultExpiration, gcInterval)

	expiration, _ := time.ParseDuration("2s")
	k1 := "hello world!"
	c.Set("k1", k1, expiration)

	if v, found := c.Get("k1"); found {
		fmt.Println("found k1:", v)
	} else {
		fmt.Println("not found k1")
	}

	err := c.SaveToFile("./items.txt")
	if err != nil {
		fmt.Println(err)
	}
	err = c.LoadFromFile("./items.txt")
	if err != nil {
		fmt.Println(err)
	}

	s, _ := time.ParseDuration("4s")
	time.Sleep(s)
	if v, found := c.Get("k1"); found {
		fmt.Println("found k1:", v)
	} else {
		fmt.Println("not found k1")
	}

}
