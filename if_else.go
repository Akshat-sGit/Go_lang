package main

import "fmt"

func main() {
	num := 10
	if num%2 == 0 { // checks if number is even
		fmt.Println("The number", num, "is even")
	} else {
		fmt.Println("The number", num, "is odd")
	}

	fmt.Println("===============================")

	// Maximum of two numbers
	num1 := 30
	num2 := 78
	if num1 > num2 {
		fmt.Println("The Maximum Number is:", num1)
	} else {
		fmt.Println("The Maximum Number is:", num2)
	}

	fmt.Println("======================================")

	// Checking number ranges
	n := 99
	if n <= 50 {
		fmt.Println(n, "is less than or equal to 50")
	} else if n <= 100 {
		fmt.Println(n, "is between 51 and 100")
	} else {
		fmt.Println(n, "is greater than 100")
	}
}
