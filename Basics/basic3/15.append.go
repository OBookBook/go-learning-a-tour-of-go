/*
Appending to a slice
https://go-tour-jp.appspot.com/moretypes/15

スライスへ新しい要素を追加
func append(s []T, vs ...T) []T

実行コマンド
go run Basics/basic3/15.append.go
*/

package main

import "fmt"

func main() {
	var s []int
	printSlice(s) // 出力結果 len=0 cap=0 []

	s = append(s, 0)
	printSlice(s) // 出力結果 len=1 cap=1 [0]

	s = append(s, 1)
	printSlice(s) // 出力結果 len=2 cap=2 [0 1]

	s = append(s, 2, 3, 4)
	printSlice(s) // 出力結果 len=5 cap=6 [0 1 2 3 4]
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
