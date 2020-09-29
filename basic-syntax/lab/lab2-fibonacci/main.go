package main

import "fmt"

func fibonacci(n int) int {
    a := 0
    b := 1
    for i := 0; i < n; i++ {
        temp := a
        a = b
        b = temp + a
    }
    return a
}

func main() {
	//Tính fibonacci từ 0 -> 15
    for n := 0; n < 15; n++ {
        result := fibonacci(n)
        fmt.Printf("Fibonacci %v = %v\n", n, result)
    }
}
