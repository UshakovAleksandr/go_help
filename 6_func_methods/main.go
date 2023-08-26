package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"sort"
	"time"
)

func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
func executeSum() {
	sum(5, 7)
	sum(3, 2, 1)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

func lambdaFunc1() {
	func() {
		fmt.Println("Hello ")
	}()
	sayWorld := func() {
		fmt.Println("World!")
	}
	sayWorld()
}

func lambdaFunc2() {
	a := func(param int) int {
		return param + 20
	}
	fmt.Println(a(29))
}

func lambdaSorting() {
	people := []string{"Alice", "Bob", "Dave"}

	sort.Slice(people, func(i, j int) bool {
		return len(people[i]) < len(people[j])
	})

	fmt.Println(people)
}

/////////////////////////////////////

func closure() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func forClosure() {
	nextInt := closure()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := closure()
	fmt.Println(newInts())

	fmt.Println(nextInt())
	fmt.Println(newInts())
}

/////////////////////////////////
type handler func(w http.ResponseWriter, r *http.Request)

func timeCount(next handler) handler {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		end := time.Now()
		timeDelta := end.Sub(start)
		log.Printf("время исполнения запроса %v\n", timeDelta.Nanoseconds())
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello!</h1>")
}

func server() {
	http.HandleFunc("/hello", timeCount(hello))
	http.ListenAndServe(":3000", nil)
}

////////////////////////////

func timeCount1(myFunc func(param int) string, param int) {
	start := time.Now()
	myFunc(param)
	end := time.Now()
	timeDelta := end.Sub(start)
	log.Printf("время исполнения запроса %v\n", timeDelta.Nanoseconds())
}

func foo(param int) string {
	for i := 0; i < param; i++ {
	}
	return "1111"
}

////////////////////////////

type SumFunc func(arguments ...int) int

func bar() {
	var summer SumFunc
	summer = func(arguments ...int) int {
		var total int
		for _, v := range arguments {
			total += v
		}
		return total
	}
	result := summer(1, 2, 3, 4)
	fmt.Println(result)
}

////////////////////////

func bar1() {
	s := "hello"
	defer fmt.Println(s)
	defer func() {
		fmt.Println(s)
	}()
	s = "world"
}

// сигнатура ответа!!!!!
func doSmth() (s int, err error) {
	defer func() {
		if err != nil {
			s = 0
		}
	}()
	s = 1
	err = errors.New("test")
	// тут может быть много логики
	// и несколько выходов из функции с ошибкой
	return
}

func main() {

}
