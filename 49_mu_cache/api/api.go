package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"49_mu_cache/db"
)

// App - приложения.
type App struct {
	router *mux.Router // маршрутизатор запросов
	db     db.Storage  // база данных
}

// New - Конструктор App.
func New(storage db.Storage) *App {
	app := App{
		router: mux.NewRouter(),
		db:     storage,
	}
	app.endpoints()
	return &app
}

// Router - возвращает маршрутизатор запросов.
func (a *App) Router() *mux.Router {
	return a.router
}

// Регистрация методов App в маршрутизаторе запросов.
func (a *App) endpoints() {
	a.router.HandleFunc("/orders", a.ordersHandler).Methods(http.MethodGet)
	a.router.HandleFunc("/orders", a.newOrderHandler).Methods(http.MethodPost)
	a.router.HandleFunc("/orders/{id}", a.updateOrderHandler).Methods(http.MethodPut)
	a.router.HandleFunc("/orders/{id}", a.deleteOrderHandler).Methods(http.MethodDelete)
}

// ordersHandler - возвращает все заказы.
func (a *App) ordersHandler(w http.ResponseWriter, r *http.Request) {
	// Получение данных из БД.
	orders := a.db.Orders()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}

// newOrderHandler - создает новый заказ.
func (a *App) newOrderHandler(w http.ResponseWriter, r *http.Request) {
	var o db.Order
	err := json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	id := a.db.NewOrder(o)
	w.Write([]byte(strconv.Itoa(id)))
}

// updateOrderHandler - обновляет данные заказа по ID.
func (a *App) updateOrderHandler(w http.ResponseWriter, r *http.Request) {
	// Считывание параметра {id} из пути запроса.
	// Например, /orders/45.
	s := mux.Vars(r)["id"]
	id, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Декодирование в переменную тела запроса,
	// которое должно содержать JSON-представление
	// обновляемого объекта.
	var o db.Order
	err = json.NewDecoder(r.Body).Decode(&o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	o.ID = id
	// Обновление данных в БД.
	a.db.UpdateOrder(o)
	// Отправка клиенту статуса успешного выполнения запроса
	w.WriteHeader(http.StatusOK)
}

// deleteOrderHandler - удаляет заказ по ID.
func (a *App) deleteOrderHandler(w http.ResponseWriter, r *http.Request) {
	s := mux.Vars(r)["id"]
	id, err := strconv.Atoi(s)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	a.db.DeleteOrder(id)
	w.WriteHeader(http.StatusOK)
}
