package main

import (
	"context"
	"fmt"
	"math"

	"github.com/redis/go-redis/v9"
)

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}

	k, err := keys(context.Background(), rdb, "1111*")
	fmt.Println(k, len(k), err)

	k, err = scanKeys(context.Background(), rdb, "1111*")
	fmt.Println(k, len(k), err)

	k, err = iterateKeys(context.Background(), rdb, "1111*")
	fmt.Println(k, len(k), err)
}

func keys(ctx context.Context, rdb *redis.Client, pattern string) ([]string, error) {
	return rdb.Keys(ctx, pattern).Result()
}

func scanKeys(ctx context.Context, rdb *redis.Client, pattern string) ([]string, error) {
	keys, _, err := rdb.Scan(ctx, 0, pattern, math.MaxInt64).Result()
	return keys, err
}

func iterateKeys(ctx context.Context, rdb *redis.Client, pattern string) ([]string, error) {
	var keys []string
	iter := rdb.Scan(ctx, 0, pattern, 0).Iterator()
	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}

	if err := iter.Err(); err != nil {
		return nil, err
	}

	return keys, nil
}

func addKeys(ctx context.Context, rdb *redis.Client) ([]string, error) {
	for i := 0; i < 30000; i++ {
		if err := rdb.Set(ctx, fmt.Sprint(i), fmt.Sprint(i), 0).Err(); err != nil {
			return nil, err
		}
	}

	keys, _, err := rdb.Scan(ctx, 0, "1*", math.MaxInt64).Result()
	return keys, err
}
