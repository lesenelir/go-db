package operation

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func SetOperation(rdb *redis.Client, ctx context.Context) {
	// set
	err := rdb.SAdd(ctx, "ai-space:go-redis:set-key1", "value1", "value2", "value3").Err()
	if err != nil {
		fmt.Println("SAdd Error: ", err)
		return
	}
	err = rdb.SAdd(ctx, "ai-space:go-redis:set-key2", "value1", "value2", "value3").Err()
	if err != nil {
		fmt.Println("SAdd Error: ", err)
		return
	}

	// get all members
	vals := rdb.SMembers(ctx, "ai-space:go-redis:set-key1").Val()
	fmt.Println(vals)

	// remove members
	err = rdb.SRem(ctx, "ai-space:go-redis:set-key1", "value2", "value3").Err()
	if err != nil {
		fmt.Println("SRem Error: ", err)
		return
	}

	// intersection
	valsIntersection, err := rdb.SInter(ctx, "ai-space:go-redis:set-key1", "ai-space:go-redis:set-key2").Result()
	if err != nil {
		fmt.Println("SInter Error: ", err)
		return
	}
	fmt.Println(valsIntersection)

	// union
	valsUnion, err := rdb.SUnion(ctx, "ai-space:go-redis:set-key1", "ai-space:go-redis:set-key2").Result()
	if err != nil {
		fmt.Println("SUnion Error: ", err)
		return
	}
	fmt.Println(valsUnion)

	// difference
	diffVals, err := rdb.SDiff(ctx, "ai-space:go-redis:set-key1", "ai-space:go-redis:set-key2").Result()
	if err != nil {
		fmt.Println("SDiff Error: ", err)
		return
	}
	fmt.Println(diffVals)
}
