package main

import (
	"fmt"
	"time"

	"github.com/schollz/progressbar/v3"
)

type WrappedFn = func(stop int) int

func TimeCount(next WrappedFn) WrappedFn {
	return func(stop int) int {
		start := time.Now()
		res := next(stop)
		fmt.Println("\n", time.Now().Sub(start).Seconds())

		return res
	}
}

func TotalCount(stop int) int {
	bar := NewBar(stop)
	var total int

	for i := 0; i < stop; i++ {
		total += i
		time.Sleep(time.Millisecond * 50)
		bar.Add(1)
	}

	return total
}

func main() {
	TotalCount := TimeCount(TotalCount)
	res := TotalCount(150)
	fmt.Println(res)
}

func NewBar(max int) *progressbar.ProgressBar {
	return progressbar.NewOptions(max,
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionSetWidth(30),
		progressbar.OptionShowIts(),
		progressbar.OptionShowBytes(true),
		//progressbar.OptionClearOnFinish(),
		//progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
		progressbar.OptionSetDescription("[cyan]counting total...[reset] "),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[green]=[reset]",
			SaucerHead:    "[green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}))
}


//package main
//
//import (
//"bytes"
//"context"
//"fmt"
//"log"
//"math"
//"os"
//"sort"
//"time"
//
//"github.com/schollz/progressbar/v3"
//"gonum.org/v1/gonum/floats"
//)
//
//var (
//	roundVal uint = 6
//	cycleVal      = 10000
//	store         = make([]float64, 0, cycleVal)
//)

//func NewBar(max int) *progressbar.ProgressBar {
//	return progressbar.NewOptions(max,
//		progressbar.OptionEnableColorCodes(true),
//		progressbar.OptionSetWidth(30),
//		progressbar.OptionShowIts(),
//		progressbar.OptionShowBytes(true),
//		progressbar.OptionShowElapsedTimeOnFinish(),
//		//progressbar.OptionClearOnFinish(),
//		//progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
//		progressbar.OptionSetDescription("[cyan]sm-salary invoke...[reset] "),
//		progressbar.OptionSetTheme(progressbar.Theme{
//			Saucer:        "[green]=[reset]",
//			SaucerHead:    "[green]>[reset]",
//			SaucerPadding: " ",
//			BarStart:      "[",
//			BarEnd:        "]",
//		}))
//}
//
//type WrappedFn = func(JSON []byte)
//
//func TimeCount(next WrappedFn) WrappedFn {
//	return func(JSON []byte) {
//		start := time.Now()
//		next(JSON)
//		store = append(store, time.Now().Sub(start).Seconds())
//	}
//}
//
//func roundFloat(val float64, precision uint) float64 {
//	ratio := math.Pow(10, float64(precision))
//	return math.Round(val*ratio) / ratio
//}
//
//func median(data []float64) float64 {
//	dataCopy := make([]float64, len(data))
//	copy(dataCopy, data)
//
//	sort.Float64s(dataCopy)
//
//	var median float64
//	l := len(dataCopy)
//	if l == 0 {
//		return 0
//	} else if l%2 == 0 {
//		median = (dataCopy[l/2-1] + dataCopy[l/2]) / 2
//	} else {
//		median = dataCopy[l/2]
//	}
//
//	return median
//}
//
//func mean(data []float64) float64 {
//	if len(data) == 0 {
//		return 0
//	}
//	var sum float64
//	for _, d := range data {
//		sum += d
//	}
//	return sum / float64(len(data))
//}
//
//func report() {
//	minVal := roundFloat(floats.Min(store), roundVal)
//	fmt.Printf("\nmin = %f сек\n", minVal)
//
//	maxVal := roundFloat(floats.Max(store), roundVal)
//	fmt.Printf("max = %f сек\n", maxVal)
//
//	medianVal := roundFloat(median(store), roundVal)
//	fmt.Printf("median = %f сек\n", medianVal)
//
//	meanVal := roundFloat(mean(store), roundVal)
//	fmt.Printf("mean = %f сек\n", meanVal)
//
//	totalTime := roundFloat(floats.Sum(store), roundVal)
//	fmt.Printf("total_time = %f сек", totalTime)
//}
//
//func fileRead() []byte {
//	data, err := os.ReadFile("sm_salary.json")
//	if err != nil {
//		log.Fatalln(err)
//	}
//
//	return data
//}
//
//func main() {
//	JSON := fileRead()
//	bar := NewBar(cycleVal)
//	invoke := TimeCount(invoke)
//
//	for i := 0; i < cycleVal; i++ {
//		invoke(JSON)
//		bar.Add(1)
//	}
//
//	report()
//}

//func decode(JSON []byte) *request.Request {
//	req := &request.Request{}
//	if err := json.Unmarshal(JSON, req); err != nil {
//		log.Fatalln(err)
//	}
//
//	return req
//}
//
//func encode(response response.Response) {
//	resp, _ := json.Marshal(response)
//	_ = resp
//}
//
//func invoke(JSON []byte) {
//	resp, _ := pkg.Execute(context.Background(), decode(JSON))
//	_ = resp
//	encode(resp)
//}

