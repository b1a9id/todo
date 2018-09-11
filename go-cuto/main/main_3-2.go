package main

import "fmt"

type Dictionary struct {
	name string
	meaning string
}

type ReadFunc func(Dictionary) string

func main() {
	var readFunc ReadFunc
	var dict Dictionary
	readFunc = readOut2
	dict.name = "コーヒー"
	dict.meaning = "コーヒー豆から作られる黒色の飲み物"
	fmt.Println(readFunc(dict))
}

func readOut2(d Dictionary) string {
	return fmt.Sprintf("「%s」 は 「%s」 という意味です", d.name, d.meaning)
}
