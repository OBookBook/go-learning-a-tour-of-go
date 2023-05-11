/*
Slices
https://go-tour-jp.appspot.com/moretypes/7

スライスは可変長

実行コマンド
go run Basics/basic3/7.slices.go
*/

package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	// []int 可変長配列
	var s []int = primes[1:4]
	fmt.Println(s) // 出力結果 [3 5 7]
}
