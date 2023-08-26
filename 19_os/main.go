package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

// go run main.go --msg my_message
// со строкой
func ex1CLI() {
	var msg string

	flag.StringVar(&msg, "msg", "hello world", "message to print")
	flag.Parse()

	fmt.Println(msg)
}

// go run main.go --port 8080
// int
func ex2CLI() {
	port := flag.Int("port", 80, "http port")
	flag.Parse()

	fmt.Printf("Port = %v\n", *port)
}

// расширенный вариант
// go run main.go --msg Hi --verbose
// go run main.go --msg Hi
// go run main.go --verbose
func ex3CLI() {
	var msg string
	flag.StringVar(&msg, "msg", "hello world", "message to print")

	verbose := flag.Bool("verbose", false, "verbose output")

	flag.Parse()

	if *verbose {
		fmt.Println("you say:", msg)
	} else {
		fmt.Println(msg)
	}
}

// go run main.go -p 8080 -h 0.0.0.0
// lib pflag
func ex4PCLI() {
	var msg string
	pflag.StringVar(&msg, "msg", "hello world", "message to print")

	verbose := pflag.BoolP("verbose", "v", false, "verbose output")
	port := pflag.IntP("port","p", 80, "http port")
	host := pflag.StringP("host", "h", "localhost", "server host")

	pflag.Parse()

	fmt.Println(*port)
	fmt.Println(*host)

	if *verbose {
		fmt.Println("you say:", msg)
	} else {
		fmt.Println(msg)
	}
}

func ex1ENV() {
	// получение и печать всех переменных
	env := os.Environ() // слайс строк "key=value"
	fmt.Println(env) // USER=rob
	// проверка есть ли такая переменная
	user, ok := os.LookupEnv("USER")
	fmt.Println(user, ok) // rob
	// получить конкретную переменную
	fmt.Printf("USER = %v\n", os.Getenv("USER"))
	// установить значение переменной
	os.Setenv("PASSWORD", "qwe123")
	fmt.Printf("PASSWORD = %v", os.Getenv("PASSWORD"))
	//// удалить переменную наслучай тестов
	//os.Unsetenv("PASSWORD")
	//fmt.Println(os.ExpandEnv("$USER lives in ${CITY}")) // "шаблонизация"
}

func main() {
	ex1ENV()
}
