package main

import "fmt"

func main() {
	fmt.Println(Sum(2, 5))
}

func Sum(left int, right int) int {
	return left + right
}
