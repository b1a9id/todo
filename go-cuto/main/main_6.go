package main

import "fmt"

func main() {
	// int型のポインタ変数
	var pointer *int
	// int型変数
	var n = 100

	// &(アドレス演算子)を使って、nのアドレスを代入
	pointer = &n

	fmt.Println("nのアドレス:", &n)
	fmt.Println("pointerの値:", pointer)

	fmt.Println("nの値:", n)
	fmt.Println("pointerの中身:", *pointer)
}
