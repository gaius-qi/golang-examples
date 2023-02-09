package smap

import (
	"fmt"
	"sync"
)

type SMAP struct {
	*sync.Map
}

func NewSMAP() *SMAP {
	return &SMAP{&sync.Map{}}
}

func (s *SMAP) Keys() []string {
	var keys []string
	s.Range(func(key, value any) bool {
		keys = append(keys, fmt.Sprint(key))
		return true
	})

	return keys
}

func (s *SMAP) Items() []string {
	var values []string
	s.Range(func(key, value any) bool {
		values = append(values, fmt.Sprint(value))
		return true
	})

	return values
}
