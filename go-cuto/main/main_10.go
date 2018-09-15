package main

import "fmt"

func main() {
	var vector struct{
		X int
		Y int
	}

	vector.X = 2
	vector.Y = 5
	fmt.Println(vector)
}
