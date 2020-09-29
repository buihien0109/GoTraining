package main

import "fmt"

func reverseString(s string) string {
	r := make([]byte, len(s))
	for i := 0; i < len(s); i++ {
		r[i] = s[len(s)-1-i]
	}
	return string(r)
}

func main() {
	str := "Hello World!"
	fmt.Println(str)
	str = reverseString(str)
	fmt.Println(str)
}
