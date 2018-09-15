package main

import "fmt"

type LoopNum int

func main() {
	var three LoopNum = 3
	three.TimesDisplay("Hello")
}

func (n LoopNum) TimesDisplay(s string) {
	for i := 0; i < int(n); i++ {
		fmt.Println(s)
	}
}
