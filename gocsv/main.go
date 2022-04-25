package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

func main() {
	s, err := newStore("data.csv")
	if err != nil {
		panic(err)
	}

	if err := s.create(&Item{
		ID:   1,
		Name: "jack",
		Age:  18,
	}); err != nil {
		panic(err)
	}

	if err := s.create(&Item{
		ID:   2,
		Name: "lucy",
		Age:  20,
	}); err != nil {
		panic(err)
	}

	items, err := s.list()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%#v", items)
}

type Item struct {
	ID   uint   `csv:"id"`
	Name string `csv:"name"`
	Age  uint   `csv:"age"`
}

type Store struct {
	path string
}

func newStore(path string) (*Store, error) {
	s := &Store{
		path: path,
	}

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	if err := gocsv.MarshalFile([]*Item{}, file); err != nil {
		return nil, err
	}

	return s, nil
}

func (s *Store) create(items ...*Item) error {
	file, err := os.OpenFile(s.path, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := gocsv.MarshalWithoutHeaders(items, file); err != nil {
		return err
	}

	return nil
}

func (s *Store) list() ([]*Item, error) {
	file, err := os.Open(s.path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var items []*Item
	if err := gocsv.UnmarshalFile(file, &items); err != nil {
		return nil, err
	}

	return items, nil
}
