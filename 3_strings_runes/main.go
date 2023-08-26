package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"unicode/utf8"
	"unsafe"
)

func ex1() {
	// Строки в Go - это неизменяемая последовательность байтов ( byte = uint8 ) //
	// src/runtime/string.go
	type stringStruct struct {
		str unsafe.Pointer
		len int
	}
}

func ex2() {
	// если нужно включить в строку кавычки или переносы строки
	// - используем обратные кавычки
	// удобно для SQL запросов //
	s := `hello
"cruel"
'world' `
	fmt.Println(s)
}

func ex3() {
	// Что можно делать со строками ?
	s := "hello world"      // создавать
	var c byte = s[0]       // получать доступ к байту(!) в строке
	var s2 string = s[5:10] // получать подстроку (в байтах!)
	s2 = s + " again"       // склеивать
	l := len(s)             // узнавать длину в байтах

	fmt.Println(c, s2, l)
}

func ex4() {
	// преобразование int to string (замена ITOA) от 0 до 9
	a := 2
	b := string(rune('0' + a))
	fmt.Println(reflect.TypeOf(b), b)
}

func ex5() {
	// Руны - это целые числа, поэтому их можно складывать:
	s := "hello " + string('0'+3) // "hello 3"
	fmt.Println(s)
	s = "hello " + string('A'+1) // "hello B"
	fmt.Println(s)
}

func runes() {
	// Для удобной работы с Unicode и UTF-8 используем пакет unicode и unicode/utf8

	//// получить первую руну из строки и ее размер в байтах
	//DecodeRuneInString(s string) (r rune, size int)
	str1 := "hello, 世界"
	for len(str1) > 0 {
		runeValue, sizeInBytes := utf8.DecodeRuneInString(str1)
		fmt.Println(string(runeValue), sizeInBytes)
		str1 = str1[sizeInBytes:]
	}

	//// получить длинну строки в рунах
	//RuneCountInString(s string) (n int)
	str2 := "hello, 世界"
	fmt.Printf("длинна в байтах - %v\n", len(str2))
	fmt.Printf("длинна в рунах - %v\n", utf8.RuneCountInString(str2))

	//// проверить валидность строки
	//ValidString(s string) bool
	valid := "he2llo, 世界"
	fmt.Println(utf8.ValidString(valid))
	invalid := string([]byte{0xff, 0xfe, 0xfd})
	fmt.Println(utf8.ValidString(invalid))
}

func strInSlice() {
	// преобразование строки в слайс байт/рун
	s := "привет"
	bytesSl := []byte(s)
	runesSl := []rune(s)
	fmt.Printf("%v\n", bytesSl)
	fmt.Printf("%v\n", runesSl)
}

func iterOnStr() {
	// Итерация по строке по байтам
	s := "привет"
	for i := 0; i < len(s); i++ {
		b := s[i]
		fmt.Println(b)
		// i строго последоваельно
		// b имеет тип byte, uint8
	}

	fmt.Println("===========================")

	// Итерация по строке по рунам
	for i, r := range s {
		fmt.Println(i, r)
		fmt.Println(string(r))
		// i может перепрыгивать значения 1,2,4,6,9...
		// r - имеет тип rune, int32
	}
}

func stringsPackage() {
	//// проверка наличия подстроки
	//Contains(s, substr string) bool
	str1 := "hello world"
	fmt.Println(strings.Contains(str1, "llo ")) // true

	//// строка начинается с ?
	//HasPrefix(s, prefix string) bool
	fmt.Println(strings.HasPrefix(str1, "he")) // true

	//// склейка строк
	//Join(a []string, sep string) string
	str2 := []string{"foo", "bar", "baz"}
	fmt.Println(strings.Join(str2, ", "))

	//// разбиение по разделителю
	//Split(s, sep string) []string
	str3 := "foo bar baz"
	fmt.Println(strings.Split(str3, " "))
}

func stringBuild() {
	//Т.к. строки read-only, каждая склейка через + или += приводит к выделению памяти.
	//Что бы оптимизировать число аллокаций используйте strings.Builder
	str1 := "hello "
	str2 := "world"
	var b strings.Builder
	b.WriteString(str1)
	b.WriteString(str2)
	resultStr := b.String()
	fmt.Println(resultStr)

	rune1 := 'a'
	rune2 := 'b'
	var c strings.Builder
	c.WriteRune(rune1)
	c.WriteRune(rune2)
	resultRunes := c.String()
	fmt.Println(resultRunes)
	//for i := 10; i >= 1; i-- {
	//	b.WriteString("Код")
	//	b.WriteRune('ъ')
	//}
}

//Необходимо написать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы/руны, например:
//
//"a4bc2d5e" => "aaaabccddddde"
//"abcd" => "abcd"
//"3abc" => "" (некорректная строка)
//"45" => "" (некорректная строка)
//"aaa10b" => "" (некорректная строка)
//"aaa0b" => "aab"
//"" => ""
//"d\n5abc" => "d\n\n\n\n\nabc"
//Как видно из примеров, разрешено использование цифр, но не чисел.
//
//В случае, если была передана некорректная строка, функция должна возвращать ошибку. При необходимости можно выделять дополнительные функции / ошибки.

var ErrInvalidString = errors.New("invalid string")

func ConvertStrToInt(str string) (int, error) {
	result, err := strconv.Atoi(str)
	if err == nil {
		return result, err
	}

	return result, err
}

func Unpack(str string) (string, error) {
	var result string
	var tempStr string
	var counter int

	if str == "" {
		return "", nil
	}

	for i, v := range str {
		if counter > 1 {
			return "", ErrInvalidString
		}
		intCheck, err := ConvertStrToInt(string(v))
		if i == 0 && err == nil {
			return "", ErrInvalidString
		}
		if err != nil {
			tempStr = string(v)
			result += tempStr
			counter = 0
		} else {
			if intCheck == 0 {
				result = result[:utf8.RuneCountInString(result)-1]
			} else {
				for i := 0; i < intCheck-1; i++ {
					result += tempStr
				}
			}
			counter += 1
		}
	}
	return result, nil
}

func main() {
	str := "a4bc2d5e"

	result, err := Unpack(str)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(result)
}
