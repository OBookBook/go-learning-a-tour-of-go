/*
Type assertions
https://go-tour-jp.appspot.com/methods/15

型アサーション
t := i.(T)

実行コマンド
go run methods/15.type-assertions.go
*/

package main

import "fmt"

func main() {
	// 型アサーションによってインターフェース型から具体的な型への変換を試みています。
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s) // hello

	s, ok := i.(string)
	fmt.Println(s, ok) // hello true

	f, ok := i.(float64)
	fmt.Println(f, ok) // 0 false

	// iは実際には文字列型stringを保持しています。そのため、実行時にランタイムエラーであるpanicが発生します。
	f = i.(float64)
	fmt.Println(f) // panic: interface conversion: interface {} is string, not float64
}
