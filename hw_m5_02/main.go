package main

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

func work1(num int, wg *sync.WaitGroup) {
	fmt.Printf("START JOB #%v\n", num)
	var calc int
	for i := 0; i <= 10; i++ {
		calc += i
	}

	fmt.Printf("END JOB #%v: calc = %v\n", num, calc)
	wg.Done()
}

func work2(num int, wg *sync.WaitGroup) {
	fmt.Printf("START JOB #%v\n", num)
	var calc int
	for i := 0; i <= 20; i++ {
		calc += i
	}
	// time.Sleep(3 * time.Second)
	fmt.Printf("END JOB #%v: calc = %v\n", num, calc)
	wg.Done()
}

func work3(num int, wg *sync.WaitGroup) {
	fmt.Printf("START JOB #%v\n", num)
	var calc int
	for i := 0; i <= 30; i++ {
		calc += i
	}

	fmt.Printf("END JOB #%v: calc = %v\n", num, calc)
	wg.Done()
}

func Monitoring(t *time.Timer, q chan int, ans chan error) {
	for {
		select {
		case <-t.C:
			fmt.Println("Время вышло")
			ans <- errors.New("time is over")
			return

		case <-q:
			fmt.Println("Все горутины отработали за две секунды")
			ans <- nil
			return

		}
	}
}

func foo() error {
	wg := &sync.WaitGroup{}
	quit := make(chan int, 1)
	ans := make(chan error)
	timer := time.NewTimer(2 * time.Second)

	go work1(1, wg)
	wg.Add(1)

	go work2(2, wg)
	wg.Add(1)

	go work3(3, wg)
	wg.Add(1)

	go Monitoring(timer, quit, ans)

	wg.Wait()
	quit <- 1

	return <-ans

}

func main() {
	err := foo()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("All goroutines were completed")
	}
}
