package main

import (
	"fmt"

	"7_interfaces_part_1/cmd"
)

type General struct {
	cmd.Store
	cmd.Worker
	cmd.Scheduler
}

func newGeneral() *General {
	return &General{
		Store:     cmd.NewStore(),
		Worker:    cmd.NewWorker(),
		Scheduler: cmd.NewScheduler(),
	}
}

func main() {
	a := newGeneral()

	a.Store.PrintStore()
	a.Worker.PrintWorker()
	a.Scheduler.PrintScheduler()

	square := cmd.NewSquare(111)
	circle := cmd.NewCircle(33)

	fmt.Println(cmd.ShapeArea(square))
	fmt.Println(cmd.ShapeArea(circle))

	slOfShapes := []cmd.Shape{square, circle}

	for _, shape := range slOfShapes {
		fmt.Println(shape.Area())
	}

	for _, shape := range slOfShapes {
		fmt.Println(cmd.ShapeArea(shape))
	}

	c := cmd.Store(&cmd.StoreStruct{})
	c.PrintStore()

	var d cmd.Store = &cmd.StoreStruct{}
	d.PrintStore()

	var r cmd.Runner
	//cmd.TypeAssertion(r) // значения nil nil

	r = &cmd.Human{Name: "alex"}
	//cmd.TypeAssertion(r) // Type: *cmd.Human Value: &cmd.Human{}
	cmd.TypeAssertionSwitch(r)

	r = &cmd.Duck{Name: "aaa"} // Type: *cmd.Duck Value: &cmd.Duck{Name:"aaa"}
	//cmd.TypeAssertion(r)
	cmd.TypeAssertionSwitch(r)

	cmd.TypeAssertionSwitch(111)
	cmd.TypeAssertionSwitch("111")
	cmd.TypeAssertionSwitch(true)
}
