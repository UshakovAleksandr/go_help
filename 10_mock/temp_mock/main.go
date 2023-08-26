package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"temp_mock/store"
)

type UseCases struct {
	Store store.Store
}

func (u *UseCases) Handler(w http.ResponseWriter, r *http.Request) {
	result := u.Store.GetUser(11)
	w.Write([]byte(result))
}

func main() {
	r := mux.NewRouter()

	use := UseCases{Store: store.NewStoreMock()}

	r.HandleFunc("/", use.Handler)

	log.Fatal(http.ListenAndServe(":8000", r))
}
