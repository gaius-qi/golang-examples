package main

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
)

func BenchmarkScanKeys(b *testing.B) {
	for i := 0; i < b.N; i++ {
		scanKeys(context.Background(), redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}), "11*")
	}
}

func BenchmarkIterateKeys(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iterateKeys(context.Background(), redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}), "11*")
	}
}

func BenchmarkKeys(b *testing.B) {
	for i := 0; i < b.N; i++ {
		keys(context.Background(), redis.NewClient(&redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		}), "11*")
	}
}
