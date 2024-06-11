package operation

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func ListOperation(rdb *redis.Client, ctx context.Context) {
	// del
	err := rdb.Del(ctx, "ai-space:go-redis:list-key1").Err()
	if err != nil {
		fmt.Println("Del Error: ", err)
		return
	}

	// lpush
	err = rdb.LPush(ctx, "ai-space:go-redis:list-key1", "list-value1", "list-value2", "list-value3").Err()
	if err != nil {
		fmt.Println("LPush Error: ", err)
		return
	}

	// rpush
	err = rdb.RPush(ctx, "ai-space:go-redis:list-key1", "list-value4", "list-value5", "list-value6").Err()
	if err != nil {
		fmt.Println("RPush Error: ", err)
		return
	}

	// get list length
	listLen, err := rdb.LLen(ctx, "ai-space:go-redis:list-key1").Result()
	if err != nil {
		fmt.Println("LLen Error: ", err)
		return
	}
	fmt.Println("LLen: ", listLen)

	// lrange
	vals := rdb.LRange(ctx, "ai-space:go-redis:list-key1", 0, -1).Val()
	fmt.Println("LRange: ", vals)

	// lpop
	lVal := rdb.LPop(ctx, "ai-space:go-redis:list-key1").Val()
	fmt.Println(lVal)

	// rpop
	rVal := rdb.RPop(ctx, "ai-space:go-redis:list-key1").Val()
	fmt.Println(rVal)

}
