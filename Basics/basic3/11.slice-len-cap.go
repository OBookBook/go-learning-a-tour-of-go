/*
Slice length and capacity
https://go-tour-jp.appspot.com/moretypes/11

スライスは長さ( length )と容量( capacity )の両方を持っています。
	len関数は、スライスまたは配列内の要素の数、つまり長さを返します。
	cap関数は、スライスまたは配列が保持できる要素の最大数、つまり容量を返します。

実行コマンド
go run Basics/basic3/11.slice-len-cap.go
*/

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s) // 実行結果 len=6 cap=6 [2 3 5 7 11 13]

	// 0から0未満
	s = s[:0]
	printSlice(s) // 実行結果 len=0 cap=6 []

	// 0から4未満
	s = s[:4]
	printSlice(s) // 実行結果 len=4 cap=6 [2 3 5 7]

	// s = [2 3 5 7] より 2から最後の要素
	s = s[2:]
	printSlice(s) // 実行結果 len=2 cap=4 [5 7]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
