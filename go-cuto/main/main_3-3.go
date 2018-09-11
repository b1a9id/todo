package main

import "fmt"

type Score2 int

func (s Score2) Show()  {
	fmt.Printf("点数は%d点です\n", s)
}

func main() {
	var myScore Score2 = 100
	myScore.Show()
}
