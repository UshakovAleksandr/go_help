package cmd

import "fmt"

type Worker interface {
	PrintWorker()
}

type WorkerStruct struct{}

func NewWorker() Worker {
	return &WorkerStruct{}
}

func (w *WorkerStruct) PrintWorker() {
	fmt.Println("This is WorkerStruct method")
}
