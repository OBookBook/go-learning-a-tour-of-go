/*
Slice literals
https://go-tour-jp.appspot.com/moretypes/9

スライスの要素を変更すると、その元となる配列の対応する要素が変更されます。

実行コマンド
go run Basics/basic3/9.slice-literals.go
*/

package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q) // 出力結果 [2 3 5 7 11 13]

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r) // 出力結果 [true false true true false true]

	// スライスの要素型が定義されています。
	s := []struct {
		i int
		b bool
	}{
		// フィールドの値を指定
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s) // 出力結果 [{2 true} {3 false} {5 true} {7 true} {11 false} {13 true}]
}
