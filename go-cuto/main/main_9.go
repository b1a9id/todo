package main

import "fmt"

func main() {
ForLabel:
	for i := 0; i < 10; i++ {
		switch  {
		case i == 3:
			break ForLabel
		default:
			fmt.Println(i)
		}
	}

	for i := 0; i < 10; i++ {
		if i == 2 {
			goto Label
		}
	}

Label:
	fmt.Println("done")
}
