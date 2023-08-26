package main

import "fmt"

// [V int | float64] - constrained
// [V int | float64] - явно указываем типы, используемые в сигнатурном интерфейсе
func ex1[V int | float64](numbers []V) V {
	var sum V

	for _, v := range numbers {
		sum += v
	}

	return sum
}

// comparable - интерфейс, который реализуют comparable типы данных:
// это - все, кроме:
// map, slice,
// struct, где полями используются map и slice
func ex2[T comparable](elems []T, search T) bool {
	for _, elem := range elems {
		if elem == search {
			return true
		}
	}

	return false
}

// [T any] == interface{}
func ex3[T any](elems ...T) {
	fmt.Println(elems)
}

// Number - кастомный интерфейс, указываемый в качестве constrained в ex4.
type Number interface {
	int | float64
}

func ex4[V Number](numbers []V) V {
	var sum V

	for _, v := range numbers {
		sum += v
	}

	return sum
}

func IndexOf[T comparable](collection []T, el T) int {
	for i, x := range collection {
		if x == el {
			return i
		}
	}
	return -1
}

func main() {
	// ex1
	intSl1 := []int{1, 2, 3, 4}
	floatSl1 := []float64{1.0, 2.0, 3.0, 4.0}
	res1 := ex1(intSl1)
	res2 := ex1[int](intSl1)
	res3 := ex1(floatSl1)
	res4 := ex1[float64](floatSl1)
	fmt.Println(res1, res2, res3, res4)

	// ex2
	intSl2 := []int{1, 2, 3, 4}
	strSl2 := []string{"1", "2", "3", "4"}
	res5 := ex2(intSl2, 1)
	res6 := ex2(strSl2, "11")
	fmt.Println(res5, res6)

	// ex3
	intSl3 := []int{1, 2, 3, 4}
	floatSl3 := []float64{1.0, 2.0, 3.0, 4.0}
	strSl3 := []string{"1", "2", "3", "4"}
	ex3(intSl3)
	ex3(floatSl3)
	ex3(strSl3)

	// ex4
	intSl4 := []int{1, 2, 3, 4}
	floatSl4 := []float64{1.0, 2.0, 3.0, 4.0}
	res7 := ex1(intSl4)
	res8 := ex1[int](intSl4)
	res9 := ex1(floatSl4)
	res10 := ex1[float64](floatSl4)
	fmt.Println(res7, res8, res9, res10)
}
