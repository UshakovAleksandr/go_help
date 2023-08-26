package db

import (
	"sync"
)

// Order - Заказ на доставку товаров.
type Order struct {
	ID     int  `json:"id"`      // номер заказа
	IsOpen bool `json:"is_open"` // открыт/закрыт
	//DeliveryTime    int64     `json:"delivery_time"`    // срок доставки
	DeliveryAddress string    `json:"delivery_address"` // адрес доставки
	Products        []Product `json:"products"`         // состав заказа
}

// Product - Товар.
type Product struct {
	ID    int     `json:"id"`    // артикул товара
	Name  string  `json:"name"`  // название
	Price float64 `json:"price"` // цена
}

type Storage interface {
	Orders() []Order
	NewOrder(o Order) int
	UpdateOrder(o Order)
	DeleteOrder(id int)
}

// DB - База данных заказов.
type DB struct {
	//id    int           // текущее значение ID для нового заказа
	mu    sync.Mutex    // мьютекс для синхронизации доступа
	store map[int]Order // БД заказов
}

// New - Конструктор БД.
func New() Storage {
	db := DB{
		//id:    1, // первый номер заказа
		store: map[int]Order{},
	}
	return &db
}

// Orders - возвращает все заказы.
func (db *DB) Orders() []Order {
	db.mu.Lock()
	defer db.mu.Unlock()

	var data []Order
	for _, v := range db.store {
		data = append(data, v)
	}
	return data
}

// NewOrder - создает новый заказ.
func (db *DB) NewOrder(o Order) int {
	db.mu.Lock()
	defer db.mu.Unlock()

	//db.id = o.ID
	db.store[o.ID] = o
	//db.id++

	return o.ID
}

// UpdateOrder - обновляет данные заказа по ID.
func (db *DB) UpdateOrder(o Order) {
	db.mu.Lock()
	defer db.mu.Unlock()

	db.store[o.ID] = o
}

// DeleteOrder - удаляет заказ по ID.
func (db *DB) DeleteOrder(id int) {
	db.mu.Lock()
	defer db.mu.Unlock()

	delete(db.store, id)
}
