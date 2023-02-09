package smap

import (
	"fmt"
	"testing"
)

func BenchmarkSMAP_Load(b *testing.B) {
	s := NewSMAP()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s.Store(fmt.Sprint(n), fmt.Sprint(n))
		b.StartTimer()
		s.Load(fmt.Sprint(n))
	}
}

func BenchmarkSMAP_Store(b *testing.B) {
	s := NewSMAP()
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		s.Store(fmt.Sprint(n), fmt.Sprint(n))
	}
}

func BenchmarkSMAP_Items(b *testing.B) {
	s := NewSMAP()
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s.Store(fmt.Sprint(n), fmt.Sprint(n))
		b.StartTimer()
		s.Items()
	}
}

func BenchmarkSMAP_Keys(b *testing.B) {
	s := NewSMAP()
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s.Store(fmt.Sprint(n), fmt.Sprint(n))
		b.StartTimer()
		s.Keys()
	}
}
