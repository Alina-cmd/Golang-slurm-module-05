package main

import (
	"fmt"
	"sync"
)

func oddNumbers(ch1 chan<- int, wg *sync.WaitGroup) {
	fmt.Println("START oddNumbers()")
	a := [5]int{1, 3, 5, 7, 9}
	count := 0
	for count < 5 {
		ch1 <- a[count]
		count++
	}
	wg.Done()
}

func evenNumbers(ch2 chan<- int, wg *sync.WaitGroup) {
	fmt.Println("START evenNumbers()")
	b := [5]int{2, 4, 6, 8, 10}
	count := 0
	for count < 5 {
		ch2 <- b[count]
		count++
	}
	wg.Done()
}

func main() {
	wg := &sync.WaitGroup{}
	var c [10]int
	ch1 := make(chan int)
	ch2 := make(chan int)

	go oddNumbers(ch1, wg)
	wg.Add(1)

	go evenNumbers(ch2, wg)
	wg.Add(1)

	count := 0
	for count < 10 {
		c[count] = <-ch1
		count++
		c[count] = <-ch2
		count++

	}

	wg.Wait()
	fmt.Println(c)

}
