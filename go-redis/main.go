package main

import (
	"context"
	"fmt"
	"math"

	"github.com/redis/go-redis/v9"
)

type User struct {
	Name string
}

func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	if _, err := rdb.Ping(context.Background()).Result(); err != nil {
		panic(err)
	}

	rdb.HSet(ctx, "user:1", "name", "Alice", "age", 30)
	rdb.HSet(ctx, "user:2", "name", "Bob", "age", 25)

	rdb.HSet(ctx, "order:100", "user_id", 1, "product", "Laptop", "amount", 1200)
	rdb.HSet(ctx, "order:101", "user_id", 2, "product", "Phone", "amount", 800)

	rdb.SAdd(ctx, "user:1:orders", 100)
	rdb.SAdd(ctx, "user:2:orders", 101)

	user, err := rdb.HGetAll(ctx, "user:1").Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(user)

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
