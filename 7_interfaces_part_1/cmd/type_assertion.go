package cmd

import "fmt"

type Runner interface {
	Run()
}

type Human struct {
	Name string
}

func (h *Human) Run() {
	fmt.Println("human runs")
}

type Duck struct {
	Name string
}

func (d *Duck) Run() {
	fmt.Println("duck runs")
}

func TypeAssertion(runner Runner) {
	fmt.Println("TypeAssertion func")
	fmt.Printf("Type: %T Value: %#v\n", runner, runner)

	if human, ok := runner.(*Human); ok {
		fmt.Printf("Type: %T Value: %#v\n", human, human)
	} else {
		fmt.Println("no type Human")
	}
}

func TypeAssertionSwitch(runner interface{}) {
	fmt.Println("TypeAssertionSwitch")
	switch v := runner.(type) {
	case *Human:
		v.Run()
	case *Duck:
		v.Run()
	case int:
		fmt.Println("это int")
	case string:
		fmt.Println("это string")
	default:
		fmt.Println("no valuable type")
	}
}
