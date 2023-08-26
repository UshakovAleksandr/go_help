package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func logic() string {
	fmt.Println("logic")
	return "string"
}

func logicCall(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	result := logic()
	w.Write([]byte(result))
}

func timeHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	time.Sleep(time.Second)
	w.Write([]byte("timeCount"))
}

func middleWare(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Println("before")
		next(w, r, ps)
		fmt.Println("after")
	}
}

func timeCount(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		start := time.Now()
		next(w, r, ps)
		fmt.Println(time.Now().Sub(start).Milliseconds())
	}
}

func main() {
	router := httprouter.New()
	router.GET("/", middleWare(logicCall))
	router.GET("/time", timeCount(middleWare(timeHandler)))

	log.Fatal(http.ListenAndServe(":8000", router))
}

type WrappedFn = func(stop int) int

func Wrapper(fn WrappedFn) WrappedFn {
	return func(stop int) int {
		start := time.Now()
		res := fn(stop)
		fmt.Println(time.Now().Sub(start).Seconds())

		return res
	}
}

func Foo(stop int) int {
	var total int

	for i := 0; i < stop; i++ {
		total += i
	}

	return total
}

//func main() {
//	Foo := Wrapper(Foo)
//	res := Foo(1000000000)
//	fmt.Println(res)
//}
