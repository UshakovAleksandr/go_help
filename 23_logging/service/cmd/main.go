package main

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"

	"service/internal/user"
	l "service/pkg/logger"
)

func main() {
	router := httprouter.New()
	logger := l.NewAllLogger()
	handler := user.New(logger)
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
	}

	log.Fatalln(server.Serve(listener))
}
