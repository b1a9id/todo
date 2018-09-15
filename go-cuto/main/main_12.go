package main

import "fmt"

func main() {
	result, remainder := Div(19, 4)
	fmt.Printf("19を4で割ると%dあまり%dです。", result, remainder)
}

func Div(left int, right int) (int, int) {
	return left / right, left % right
}
