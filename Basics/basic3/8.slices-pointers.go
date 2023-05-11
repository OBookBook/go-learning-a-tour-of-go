/*
Slices are like references to arrays
https://go-tour-jp.appspot.com/moretypes/8

スライスの要素を変更すると、その元となる配列の対応する要素が変更されます。

実行コマンド
go run Basics/basic3/8.slices-pointers.go
*/

package main

import "fmt"

func main() {
	names := [4]string{
		"john",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names) // 出力結果 [john Paul George Ringo]

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b) // 出力結果 [john Paul] [Paul George]

	// スライスの要素を変更すると、その元となる配列の対応する要素 Paul が XXX] に 変更されます。
	b[0] = "XXX"
	fmt.Println(a, b)  // 出力結果 [john XXX] [XXX George]
	fmt.Println(names) // 出力結果 [john XXX George Ringo]
}
