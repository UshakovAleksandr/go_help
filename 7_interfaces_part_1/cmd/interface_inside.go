package cmd

//type iface struct {
//	tab  *itab // Информация об интерфейсе
//	data unsafe.Pointer // Хранимые данные
//}
//// itab содержит тип интерфейса и информацию о хранимом типе.
//type itab struct {
//	inter *interfacetype // Метаданные интерфейса
//	_type *_type // Go-шный тип хранимого интерфейсом значения
//	hash  uint32
//	_     [4]byte
//	fun   [1]uintptr // Список методов типа, удовлетворяющих интерфейсу
//}

// Создание интерфейса:
// - аллокация места для хранения адреса ресивера
// - получение itab:
//      - проверка кэша
//      - нахождение реализаций методов
// - создание iface: runtime.convT2I
//s := Speaker(Human{Greeting: "Hello"})
//// Динамический диспатчинг
//// - для рантайма это вызов n-го метода s.Method_0()
//// - превращается в вызов вида s.itab.fun[0](s.data)
//s.SayHello()

//type Addifier interface {
//	Add(a, b int32) int32
//}
//type Adder struct{ id int32 }
//
//func (adder Adder) Add(a, b int32) int32 {
//	return a + b
//}
//func BenchmarkDirect(b *testing.B) {
//	adder := Adder{id: 6754}
//	for i := 0; i < b.N; i++ {
//		adder.Add(10, 32)
//	}
//}
//func BenchmarkInterface(b *testing.B) {
//	adder := Addifier(Adder{id: 6754})
//	for i := 0; i < b.N; i++ {
//		adder.Add(10, 32)
//	}
//}

//
//BenchmarkDirect-16      1000000000   0.2436 ns/op   0 B/op   0 allocs/op
//BenchmarkInterface-16   957668390    1.157 ns/op    0 B/op   0 allocs/op
//$ GOOS=linux GOARCH=amd64 go tool compile -m addifier.go
//Addifier(adder) escapes to heap
