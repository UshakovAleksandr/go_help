package env_ex

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/alexflint/go-arg"
)

type Config1 struct {
	Host string
	Port int
}

// SetGetEnv - запуск с set
func SetGetEnv() {
	// установить значение переменной
	os.Setenv("HOST", "0.0.0.0")
	os.Setenv("PORT", "8080")

	httpHost := os.Getenv("HOST")
	if httpHost == "" {
		log.Fatal("HOST is not define")
	}

	httpPort, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("PORT is not define")
	}

	config := Config1{Host: httpHost, Port: httpPort}
	fmt.Printf("%+v\n", config)
}

// GetEnv - set из cli
// HOST=localhost PORT=7777 go run main.go
func GetEnv() {
	httpHost := os.Getenv("HOST")
	if httpHost == "" {
		log.Fatal("HOST is not define")
	}

	httpPort, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("PORT is not define")
	}

	config := Config1{Host: httpHost, Port: httpPort}
	fmt.Printf("%+v\n", config)
}

type Config2 struct {
	Host string `arg:"env" default:"0.0.0.0"`
	Port int    `arg:"env" default:"8080"`
}

func Ex1Lib() {
	config := &Config2{}
	arg.MustParse(config)
	fmt.Println(config.Host, config.Port)
}
