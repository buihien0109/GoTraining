package main

import "fmt"


func factorial(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func main() {
	num := 4
	result := factorial(num)
	fmt.Printf("%v! = %v\n", num, result)
}
