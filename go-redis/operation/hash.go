package operation

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func HashOperation(rdb *redis.Client, ctx context.Context) {
	// hset
	err := rdb.HSet(ctx, "ai-space:go-redis:hash-key1", "field1", "hash-key-value1").Err()
	if err != nil {
		fmt.Println("HSet Error: ", err)
		return
	}

	// hget
	val, err := rdb.HGet(ctx, "ai-space:go-redis:hash-key1", "field1").Result()
	if err != nil {
		fmt.Println("HGet Error: ", err)
		return
	}
	println("HGet: ", val)

	// hset multiple, === hmset
	err = rdb.HSet(ctx, "ai-space:go-redis:hash-key2", "field1", "value1", "field2", "value2", "field3", "value3", "field4", "value4").Err()
	if err != nil {
		fmt.Println("HSet Error: ", err)
		return
	}

	// hmget
	vals, err := rdb.HMGet(ctx, "ai-space:go-redis:hash-key2", "field1", "field2", "field3", "field4").Result()
	if err != nil {
		fmt.Println("HMGet Error: ", err)
		return
	}
	for i, v := range vals { // []interface{}
		fmt.Printf("HMGet %d: %v\n", i, v)
	}

	// hgetall
	vals2, err := rdb.HGetAll(ctx, "ai-space:go-redis:hash-key2").Result()
	if err != nil {
		fmt.Println("HGetAll Error: ", err)
		return
	}
	// vals2 is map[string]string
	for key, value := range vals2 {
		fmt.Printf("HGetAll %s: %s\n", key, value)
	}

	// hdel : delete fields
	err = rdb.HDel(ctx, "ai-space:go-redis:hash-key2", "field3", "field4").Err()
	if err != nil {
		fmt.Println("HDel Error: ", err)
		return
	}

}
