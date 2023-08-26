package main

func OldFillSlice() []int {
	const size = 32000

	a := make([]int, 0)
	for i := 0; i < size; i++ {
		a = append(a, i)
	}
	return a
}

func NewFillSlice() []int {
	const size = 32000

	a := make([]int, 0, size)
	for i := 0; i < size; i++ {
		a = append(a, i)
	}
	return a
}

// Запустить конкретный бенчмарк
// go test -bench=BenchmarkFillSlice
// BenchmarkFillSlice-16   7722(кол-во запусков) 153646 ns/op(время на 1 операцию)

// Бенчить использование памяти
// go test -bench=BenchmarkFillSlice -benchmem
// go test -bench=. -benchmem
// 6907(кол-во запусков)  151532 ns/op(время на 1 операцию)
// 1418492 B/op(сколь байт алл на операцию) 25 allocs/op(сколько алл было куче, не на стеке!!!)

// Бенчить 10 секунд
//  go test -bench=BenchmarkFillSlice -benchmem -benchtime=10s -count=5

// Сравнить два бенчмарка
// go get golang.org/x/perf/cmd/benchstat
// go get golang.org/x/tools/cmd/benchcmp

// go test -bench=BenchmarkFillSlice -benchmem > old.out
// go test -bench=BenchmarkFillSlice -benchmem > new.out

// benchstat -delta-test none old.out new.out
// benchcmp old.out new.out
// benchstat -delta-test none old.out new.out && benchcmp old.out new.out

// go test -bench=. -benchmem

// команда создает !!!3!!! файла
// go test -bench=. -cpuprofile=cpu.out -memprofile=mem.out .
// go tool pprof -http=":8090" 25_bench.test cpu.out

func main() {

}

// go test . -bench . -benchmem -benchtime=15s -memprofile mem.prof
// go tool pprof -http :9000 mem.prof
