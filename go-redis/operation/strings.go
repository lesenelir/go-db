package operation

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func StringsOperation(rdb *redis.Client, ctx context.Context) {
	// set: expiration - us
	err := rdb.Set(ctx, "key1", "value1", 2_000_000).Err()
	if err != nil {
		fmt.Println("Set Error: ", err)
		return
	}

	// get
	val, err := rdb.Get(ctx, "key1").Result()
	if err != nil {
		fmt.Println("Get Error: ", err)
		return
	}
	fmt.Println("Get key1: ", val)

	// mset
	err = rdb.MSet(ctx, "ai-space:go-redis:key1", "value1", "ai-space:go-redis:key2", "value2").Err()
	if err != nil {
		fmt.Println("MSet Error: ", err)
		return
	}

	// mget
	values, err := rdb.MGet(ctx, "ai-space:go-redis:key1", "ai-space:go-redis:key2").Result()
	if err != nil {
		fmt.Println("MGet Error: ", err)
		return
	}
	fmt.Println(values)

}
