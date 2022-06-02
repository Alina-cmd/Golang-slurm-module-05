package main

import "fmt"

func foo() {
	defer func() {
		err := recover()
		fmt.Println("Восстановление")
		fmt.Println(err)
	}()
	panic("panic in foo()")
}

func main() {
	foo()
}
