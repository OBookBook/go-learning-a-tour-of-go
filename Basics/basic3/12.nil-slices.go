/*
Nil slices
https://go-tour-jp.appspot.com/moretypes/12

スライスのゼロ値は nil です。

実行コマンド
go run Basics/basic3/12.nil-slices.go
*/

package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s)) // 出力結果 [] 0 0
	if s == nil {
		fmt.Println("nil!") // 出力結果 nil!
	}
}
