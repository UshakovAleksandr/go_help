package main

import (
	"fmt"

	"22_config/configFromFile"
)

func main() {
	//env_ex.SetGetEnv()
	//env_ex.GetEnv() // HOST=localhost PORT=7777 go run main.go
	//env_ex.Ex1Lib() // HOST=localhost PORT=7777 go run main.go

	r := configFromFile.GetFileConfig() // HOST=1.1.1.1 PORT=5555 go run main.go
	fmt.Printf("HOST: %v\nPORT: %v\n", r.Host, r.Port)
}
