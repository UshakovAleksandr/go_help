package main

import (
	_ "net/http/pprof"
	"runtime"
)

type BufFreeList struct {
	ch chan []int
}

func (p *BufFreeList) Get() []int {
	select {
	case b := <-p.ch:
		return b
	default:
		return make([]int, 0)
	}
}

func (p *BufFreeList) Put(b []int) {
	select {
	case p.ch <- b[:0]: // ok
	default: // drop
	}
}

func NewBufFreeList(max int) *BufFreeList {
	c := make(chan []int, max)
	for i := 0; i < max; i++ {
		c <- make([]int, 0)
	}
	return &BufFreeList{ch: c}
}

var Cache = NewBufFreeList(runtime.NumCPU())

func Foo(num int) []int {
	//sl := Cache.Get()
	var sl []int

	if num > 10 {
		sl = append(sl, 0)
	}
	if num > 2000 {
		sl = append(sl, 1)
	}
	if num > 10000 {
		sl = append(sl, 2)
	}

	return sl
}

func start() {
	for i := 0; i < 15000; i++ {
		Foo(i)
		//Cache.Put(Foo(i))
	}
}

func main() {
	start()
}

// 1. импортируем - _ "net/http/pprof"
// 2. создаем бенчмарк на тестируемую, либо вызывающую функцию
// 3. make test-for-pprof создает 2 файла file.test и mem.prof
// 4. make pprof запускает сервер на 9000 порту для отображения mem.prof в html
// 5. viev -> top - выделяем строки, которые хотим рассмотреть детально и дальше viev -> sourse
