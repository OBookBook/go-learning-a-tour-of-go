/*
Arrays
https://go-tour-jp.appspot.com/moretypes/6

配列は固定長
[n]T 型は、型 T の n 個の変数の配列( array )を表します

実行コマンド
go run Basics/basic3/6.array.go
*/

package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1]) // 出力結果 Hello World
	fmt.Println(a)          // 出力結果 [Hello World]

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes) // 出力結果 [2 3 5 7 11 13]
}
