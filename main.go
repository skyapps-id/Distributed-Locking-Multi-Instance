package main

import (
	"context"
	"dle/redis"
	"fmt"
	"strconv"
	"sync"
)

var ctx = context.Background()
var rds = redis.NewRedisClient()

func decriment() {
	rs := rds.Redsync()
	mutex := rs.NewMutex("promo-lock")

	mutex.Lock()
	val, err := rds.Get(ctx, "PROMO")
	if err != nil {
		fmt.Println("not found value")
	}

	toInt, _ := strconv.Atoi(val.(string))
	toInt--
	if err := rds.Set(ctx, "PROMO", toInt, 0); err != nil {
		fmt.Println(err)
	}
	mutex.Unlock()
}

func main() {
	// Init Promo
	if err := rds.Set(ctx, "PROMO", 30000, 0); err != nil {
		fmt.Println(err)
	}

	// Run multi threading
	var wg sync.WaitGroup
	doIncrement := func(n int) {
		for i := 0; i < n; i++ {
			decriment()
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement(10000)
	go doIncrement(10000)
	go doIncrement(10000)
	wg.Wait()

	val, _ := rds.Get(ctx, "PROMO")
	fmt.Println(val)
	// Result 0
}
