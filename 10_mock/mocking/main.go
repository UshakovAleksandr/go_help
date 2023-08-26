package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"mocking/handler"
	"mocking/store"
)

func YourHandler(w http.ResponseWriter, r *http.Request) {
	i := store.NewStore()
	h := handler.NewHandler(i)
	res1 := h.GetUserHandler(1)
	res2 := h.CreateUserHandler("1")
	fmt.Println(res1)
	fmt.Println(res2)
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", YourHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}

//func main() {
//	//i := store.NewStore()
//	//h := handler.NewHandler(i)
//	//r := h.GetUserHandler(1)
//	//fmt.Println(r)
//}

// go get -u github.com/golang/mock/gomock
// go get github.com/golang/mock/mockgen
