package smap

import (
	"fmt"
	"testing"

	cmap "github.com/orcaman/concurrent-map/v2"
)

func BenchmarkCMAP_Get(b *testing.B) {
	c := cmap.New[string]()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		c.Set(fmt.Sprint(n), fmt.Sprint(n))
		b.StartTimer()
		c.Get(fmt.Sprint(n))
	}
}

func BenchmarkCMAP_Set(b *testing.B) {
	c := cmap.New[string]()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		c.Set(fmt.Sprint(n), fmt.Sprint(n))
	}
}

func BenchmarkCMAP_Items(b *testing.B) {
	c := cmap.New[string]()
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		c.Set(fmt.Sprint(n), fmt.Sprint(n))
		b.StartTimer()
		c.Items()
	}
}

func BenchmarkCMAP_Keys(b *testing.B) {
	c := cmap.New[string]()
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		c.Set(fmt.Sprint(n), fmt.Sprint(n))
		b.StartTimer()
		c.Keys()
	}
}
