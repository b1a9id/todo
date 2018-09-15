package main

import "fmt"

func main() {
	a, b := 10, 10
	called(a, &b)

	fmt.Println("値渡し:", a)
	fmt.Println("ポインタ渡し:", b)
}

func called(a int, b *int) {
	// 変数をそのまま変更
	a = a + 1
	// 変数の中身を変更
	*b = *b + 1
}
