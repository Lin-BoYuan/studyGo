package main

import "fmt"

func main() {
	// 没赋值则与上一行相同
	const (
		A = "AAA"
		B = "BBB"
		C
		D
	)

	fmt.Println(A, B, C, D)
}
