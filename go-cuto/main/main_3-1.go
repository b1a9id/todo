package main

import "fmt"

func main() {
	var readFunc func(struct{name string; meaning string}) string
	var dict struct{name string; meaning string}

	readFunc = readOut
	dict.name = "コーヒー"
	dict.meaning = "コーヒー豆から作られる黒色の飲み物"
	fmt.Println(readFunc(dict))
}

func readOut(s struct{name string; meaning string}) string {
	return fmt.Sprintf("「%s」は「%s」という意味です。", s.name, s.meaning)
}
