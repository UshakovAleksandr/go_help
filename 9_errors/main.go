package main

import (
	"errors"
	"fmt"
	"log"
)

func errCreate() {

	// 1 вариант
	err := errors.New("Im an error")
	if err != nil {
		fmt.Print(err)
	}

	// 2 вариант
	whoami := "error"
	err = fmt.Errorf("Im an %s", whoami)
	if err != nil {
		fmt.Print(err)
	}

}

//// Проверка ошибок: типы
//func errTypeSwitch() {
//	err := readConfig()
//	switch err := err.(type) {
//	case nil:
//		// call succeeded, nothing to do
//	case *PathError:
//		fmt.Println(“invalid config path:”, err.Path)
//	default:
//		// unknown error
//	}
//}
//
//func inrerfaceErrSwitch() {
//	type Error interface {
//		error
//		Timeout() bool   // Is the error a timeout?
//		Temporary() bool // Is the error temporary?
//	}
//
//	if nerr, ok := err.(net.Error); ok && nerr.Temporary() {
//		time.Sleep(1e9)
//		continue
//	}
//	if err != nil {
//		log.Fatal(err)
//	}
//}
//
//func ex() {
//	if err.Error() == "smth" { // Строковое представление - для людей.
//	}
//	func Write(w io.Writer, buf []byte) {
//		w.Write(buf) // Забыли проверить ошибку
//	}
//	func Write(w io.Writer, buf []byte) error {
//		_, err := w.Write(buf)
//		if err != nil {
//		// Логируем ошибку вероятно несколько раз
//		// на разных уровнях абстракции.
//		log.Println("unable to write:", err)
//		return err
//	}
//		return nil }
//}
//
//func ex2() {
//	switch err := errors.Cause(err).(type) {
//	case *MyError:
//		// handle specifically
//	default:
//		// unknown error
//	}
//
//	// IsTemporary returns true if err is temporary.
//	func IsTemporary(err error) bool {
//		te, ok := errors.Cause(err).(temporary)
//		return ok && te.Temporary()
//	}
//}

//func defer1() {
//	func Contents(filename string) (string, error) {
//		f, err := os.Open(filename)
//		if err != nil {
//			return "", err
//		}
//		defer f.Close()  // f.Close will run when we're finished.
//		var result []byte
//		buf := make([]byte, 100)
//		for {
//			n, err := f.Read(buf[0:])
//			result = append(result, buf[0:n]...)
//			if err != nil {
//				return "", err  // f will be closed if we return here.
//			}
//		}
//		return string(result), nil // f will be closed if we return here.
//	}
//}

func errRealization() {
	//func New(text string) error {
	//	return &errorString{text}
	//}
	//type errorString struct {
	//	s string
	//}
	//func (e *errorString) Error() string {
	//	return e.s
	//}
}

//func (router HttpRouter) parse(reader *bufio.Reader) (Request, error) {
//	requestText, err := readCRLFLine(reader)
//	if err != nil {
//		return nil, err
//	}
//	requestLine, err := parseRequestLine(requestText)
//	if err != nil {
//		return nil, err
//	}
//	if request := router.routeRequest(requestLine); request != nil {
//		return request, nil
//	}
//	return nil, requestLine.NotImplemented()
//}

func someFunc(i int) (int, error) {
	j, err := funcReturningError(i)
	if err != nil {
		return 0, fmt.Errorf("wrap error: %w", err)
	}

	return j, nil
}

func funcReturningError(i int) (int, error) {
	if i == 0 {
		return 0, fmt.Errorf("got %d", i)
	}

	return i, nil
}

// Аргументы отложенного вызова функции вычисляются тогда, когда вычисляется команда defer.
// 0
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

// 1
func b() {
	i := 0
	defer func() { fmt.Print(i) }()
	i++
	return
}

// 0
func c() {
	i := 0
	defer func(i int) { fmt.Print(i) }(i)
	i++
	return
}

// 3210
func d() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

// 2
func e() (i int) {
	defer func() { i++ }()
	return 1
}

//// recover
//func server(workChan <-chan *Work) {
//	for work := range workChan {
//		go safelyDo(work)
//	}
//}
//func safelyDo(work *Work) {
//	defer func() {
//		if err := recover(); err != nil {
//			log.Println("work failed:", err)
//		}
//	}()
//	do(work)
//}

func makePanic() {
	panic("aaa")
}

func panicRecover() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic name:", err)
		}
	}()

	makePanic()
}

func main() {

	panicRecover()

	//i, err := someFunc(0)
	//fmt.Println(i, err)
	//
	//i, err = someFunc(10)
	//fmt.Println(i, err)
}
