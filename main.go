package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello world")

	var age int
	fmt.Println("The age is: ", age)
	age = 29
	fmt.Println("Now the age is: ", age)
	age = 35
	fmt.Println("The age now is: ", age)
	const a = 2
	const (
		b = 23
		c = "This is a string"
	)
	fmt.Println(reflect.TypeOf(a))
}

func addAndMultiply(a, b int) (int, int) {
	return a + b, a * b
}
