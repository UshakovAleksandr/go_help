package cmd

import "fmt"

type Scheduler interface {
	PrintScheduler()
}

type SchedulerStruct struct{}

func NewScheduler() Scheduler {
	return &SchedulerStruct{}
}

func (s *SchedulerStruct) PrintScheduler() {
	fmt.Println("This is SchedulerStruct method")
}
