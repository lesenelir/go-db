package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go-redis/operation"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	defer rdb.Close()

	// ---------------------------
	// normal
	val, err := rdb.Get(ctx, "name").Result()
	if err != nil {
		fmt.Println("Get Error: ", err)
		return
	}

	fmt.Println("Get name: ", val)

	val2, err := rdb.Get(ctx, "name2").Result()
	if errors.Is(err, redis.Nil) {
		fmt.Println("name2 key does not exist")
		// no return here, so the code continues to execute
	} else if err != nil {
		fmt.Println("Get Error: ", err)
	} else {
		fmt.Println("Get name2: ", val2)
	}

	// ---------------------------
	// strings
	operation.StringsOperation(rdb, ctx)
	// hash
	operation.HashOperation(rdb, ctx)
	// list
	operation.ListOperation(rdb, ctx)
	// set
	operation.SetOperation(rdb, ctx)
	// sorted set
	operation.SortedSetOperation(rdb, ctx)
}
