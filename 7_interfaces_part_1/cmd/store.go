package cmd

import "fmt"

type Store interface {
	PrintStore()
}

type StoreStruct struct{}

func NewStore() Store {
	return &StoreStruct{}
}

func (s *StoreStruct) PrintStore() {
	fmt.Println("This is StoreStruct method")
}
