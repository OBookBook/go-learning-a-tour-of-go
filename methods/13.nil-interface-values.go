/*
Nil interface values
https://go-tour-jp.appspot.com/methods/13

実行コマンド
go run methods/13.nil-interface-values.go
*/
package main

import "fmt"

type I interface {
	M()
}

// 呼び出す 具体的な メソッドを示す型がインターフェースのタプル内に存在しないため、
// nil インターフェースのメソッドを呼び出すと、ランタイムエラーになります。
func main() {
	var i I
	describe(i) // runtime error:
	i.M()       // runtime error:
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
