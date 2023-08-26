package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func emulateLongOperation(ctx context.Context, cancel context.CancelFunc, id int, wg *sync.WaitGroup) {
	defer wg.Done()

	randomizer := rand.New(rand.NewSource(time.Now().UnixNano()))
	randVal := randomizer.Intn(200)

	timer := time.NewTimer(time.Millisecond * time.Duration(randVal))
	defer timer.Stop()

	select {
	case <-ctx.Done():
		//fmt.Println(ctx.Err())
		log.Printf("Job %d timeout", id)
	case <-timer.C:
		log.Printf("Job %d done", id)
		cancel()
	}

}

// RunJobs - без передачи аргумента cancel - выполнить jobs которые успевают за 100 мс
// c cancel - только первый (иногда выполнится несколько), точно что бы 1 - runtime.GOMAXPROCS(1)
func RunJobs(ctx context.Context, cancel context.CancelFunc) {

	defer cancel()

	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go emulateLongOperation(ctx, cancel, i, wg)
	}

	wg.Wait()
}

func doDbRequest(ctx context.Context) {
	data, ok := ctx.Value("trace_id").(int)
	if ok {
		fmt.Println(data)
	}
	fmt.Println("Request done")
}

func main() {
	//runtime.GOMAXPROCS(1)
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	RunJobs(ctx, cancel)

	traceCtx := context.WithValue(context.Background(), "trace_id", 133)
	doDbRequest(traceCtx)
}
