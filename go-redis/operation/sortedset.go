package operation

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

func SortedSetOperation(rdb *redis.Client, ctx context.Context) {
	// zadd: add members with scores
	err := rdb.ZAdd(
		ctx,
		"ai-space:go-redis:sortedset-key1",
		redis.Z{Score: 30, Member: "user1"},
		redis.Z{Score: 90, Member: "user2"},
		redis.Z{Score: 60, Member: "user3"},
		redis.Z{Score: 10, Member: "user4"},
		redis.Z{Score: 40, Member: "user5"},
		redis.Z{Score: 50, Member: "user6"},
	).Err()
	if err != nil {
		fmt.Println("ZAdd Error: ", err)
		return
	}

	// zcard: get the number of members
	num := rdb.ZCard(ctx, "ai-space:go-redis:sortedset-key1").Val()
	fmt.Println("The number of members: ", num)

	// zscore: get the score of a member
	score, err := rdb.ZScore(ctx, "ai-space:go-redis:sortedset-key1", "user1").Result()
	if err != nil {
		fmt.Println("ZScore Error: ", err)
		return
	}
	fmt.Println("The score of user1: ", score)

	// zrange: get members by index (score ascending order)
	members := rdb.ZRange(ctx, "ai-space:go-redis:sortedset-key1", 0, -1).Val()
	fmt.Println("ZRange: ", members)

	// zrevrange: get members by index (score descending order)
	membersDesc := rdb.ZRevRange(ctx, "ai-space:go-redis:sortedset-key1", 0, -1).Val()
	fmt.Println("ZRevRange: ", membersDesc)

}
