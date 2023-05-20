/*
The empty interface
https://go-tour-jp.appspot.com/methods/14

空のインターフェースは、任意の型の値を保持できます。

実行コマンド
go run methods/14.empty-interface.go
*/

package main

import "fmt"

func main() {
	var i interface{}
	describe(i) // (<nil>, <nil>)

	i = 42
	describe(i) // (42, int)

	i = "hello"
	describe(i) // (hello, string)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
