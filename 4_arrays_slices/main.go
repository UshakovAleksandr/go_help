package main

import (
	"fmt"
	"sort"
	"unsafe"
)

func arrays1() {
	var arr1 [256]int         // фиксированная длина
	var arr2 [10][10]int      // может быть многомерным
	arr3 := [...]int{1, 2, 3} // автоматический подсчет длины
	arr4 := [10]int{1, 2, 3, 4, 5}

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)
}

func arrays2() {
	var arr [10]int
	v := arr[1] // чтение
	fmt.Println(v)
	arr[3] = 1            // запись
	fmt.Println(len(arr)) // длина массива
	fmt.Println(arr[2:4]) // получение слайса
}

func slices1() {
	var a []int             // не-инициализированный слайс, nil
	b := []int{}            // с помощью литерала слайса
	c := make([]int, 3, 10) // с помощью функции make, s == {0,0,0}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(len(c), cap(c))
}

func slices2() {
	// runtime/slice.go
	type slice struct {
		array unsafe.Pointer
		len   int
		cap   int
	}

	s := make([]int, 3, 10)

	l := len(s) // len — вернуть длину слайса
	c := cap(s) // cap — вернуть емкость слайса

	fmt.Println(l, c)

}

func slices3() {
	a := []int{100, 200, 300, 1, 500}
	v := a[1] // чтение
	fmt.Println(v)

	a[3] = 400          // запись
	fmt.Println(len(a)) // длина массива

	fmt.Println(a[2:4]) // получение слайса
}

func slices4() {
	//s = append(s, 1) // добавляет 1 в конец слайса
	//s = append(s, 1, 2, 3) // добавляет 1, 2, 3 в конец слайса
	//s = append(s, s2...) // добавляет содержимое слайса s2 в конец s

	var s []int      // s == nil
	s = append(s, 1) // s == {1} append умеет работать с nil-слайсами
}

func slices5() {
	s := []int{1}
	for i := 0; i < 10; i++ {
		// слайс в памяти - указатель на первый элемент
		fmt.Printf("ptr %v   len: %d cap %d  \n", &s[0], len(s), cap(s))
		s = append(s, i)
	}
}

// s[i:j] — возвращает под-слайс, с i -ого элемента включительно, по j -ый не влючительно. Длинна нового слайса будет j-i .

func slicing1() {
	s := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s2 := s[3:5] // [3 4]
	s3 := s[3:]  // [3 4 5 6 7 8 9]
	s4 := s[:5]  // [0 1 2 3 4]
	s5 := s[:]   // копия s (shallow) [0 1 2 3 4 5 6 7 8 9]

	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
}

func slicing2() {
	s := []byte{1, 2, 3, 4, 5}
	s2 := s[2:5] // останется указатель на тот же массив, что и у s
	fmt.Println(s2)
}

func slicing3() {

	//arr := []int{1, 2}
	//arr2 := arr // копируется только заголовок, массив остался общий arr2[0] = 42
	//fmt.Println(arr[0]) // ?
	//arr2 = append(arr2, 3, 4, 5, 6, 7, 8) // реаллокация
	//arr2[0] = 1
	//fmt.Println(arr[0]) // ?

	arr := []int{1, 2}
	arr2 := arr
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println("<----------------------->")

	arr2[0] = 42
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println("<----------------------->")

	arr2 = append(arr2, 3, 4, 5, 6, 7, 8, 9, 0)
	fmt.Println(arr)
	fmt.Println(arr2)
	fmt.Println("<----------------------->")

	arr2[0] = 1
	fmt.Println(arr)
	fmt.Println(arr2)
}

// можно без return
func slices6(sl []int) {
	sl[0] = 100
}

// c return обязательно, если есть изменения размерности слайса
func slices7(sl []int) []int {
	sl = append(sl, 10)
	return sl
}

//func addValue(slice []int) []int {
//	return append(slice, 4)
//}
//
//initialSlice := []int{1, 2}
//fmt.Printf("Type: %T Value: %#v\n", initialSlice, initialSlice)
//fmt.Printf("Length: %d Capacity: %d\n\n", len(initialSlice), cap(initialSlice))
//
//newSlice := append(initialSlice, 3)
//fmt.Printf("Type: %T Value: %#v\n", newSlice, newSlice)
//fmt.Printf("Length: %d Capacity: %d\n\n", len(newSlice), cap(newSlice))
//
//newSlice2 := addValue(newSlice) // тут не превысили cap, поэтому newSlice и newSlice2 под капотом ссылаются на один и тот же массив
//fmt.Printf("Type: %T Value: %#v\n", newSlice2, newSlice2)
//fmt.Printf("Length: %d Capacity: %d\n\n", len(newSlice2), cap(newSlice2))
//
//newSlice2[0] = 500 // если меняем значение в одном слайсе, то это отражается и на втором
//
//fmt.Printf("Type: %T Value: %#v\n", newSlice, newSlice)
//fmt.Printf("Type: %T Value: %#v\n", newSlice2, newSlice2)
//
//newSlice3 := addValue(newSlice2) // тут мы превысили cap и у newSlice3 уже новый массив под капотом
//newSlice3[3] = 1000              // даже если мы поменяем что то в новом массиве через newSlice3, другие слайсы этого не увидят
//
//fmt.Printf("Type: %T Value: %#v\n", newSlice2, newSlice2)
//fmt.Printf("Type: %T Value: %#v\n", newSlice3, newSlice3)

func sort1() {
	i := []int{3, 2, 1}
	sort.Ints(i)
	fmt.Println(i)

	s := []string{"hello", "cruel", "world"}
	sort.Strings(s)
	fmt.Println(s)
}

func sort2() {
	type User struct {
		Name string
		Age  int
	}

	s := []User{
		{"vasya", 19},
		{"petya", 18},
	}
	sort.Slice(s, func(i, j int) bool {
		return s[i].Age < s[j].Age
	})
	fmt.Println(s)

}

// Напишите функцию `Concat`, которая получает несколько слайсов
// и склеивает их в один длинный.
// { {1, 2, 3}, {4, 5}, {6, 7} }  => {1, 2, 3, 4, 5, 6, 7}

//func Concat(slices [][]int) []int {
//	var result []int
//	for _, v := range slices {
//		result = append(result, v...)
//	}
//	return result
//}
//
//func TestConcat(t *testing.T) {
//	test := []struct {
//		slices   [][]int
//		expected []int
//	}{
//		{[][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
//		{[][]int{{1, 2}, {3, 4}, {6, 5}}, []int{1, 2, 3, 4, 6, 5}},
//		{[][]int{{1, 2}, {}, {6, 5}}, []int{1, 2, 6, 5}},
//	}
//
//	for _, tc := range test {
//		require.Equal(t, tc.expected, Concat(tc.slices))
//	}
//}

func main() {
	sort2()
	// https://github.com/golang/go/wiki/SliceTricks
}
