/*
Slice defaults
https://go-tour-jp.appspot.com/moretypes/10

これらのスライス式は等価です:
a[0:10] 0から10未満
a[:10] 0から10未満
a[0:] 0から最後の要素
a[:] 0から最後の要素

実行コマンド
go run Basics/basic3/10.slice-bounds.go
*/

package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	// インデックス 1 からインデックス 4 未満までの要素を取得
	s = s[1:4]
	fmt.Println(s) // 出力結果 [3 5 7]

	// s = [3 5 7] より 最初の要素からインデックス 2 未満までの要素を取得
	s = s[:2]
	fmt.Println(s) // 出力結果 [3 5]

	// s = [3 5] インデックス 1 から最後の要素までの要素
	s = s[1:]
	fmt.Println(s) // 出力結果 [5]
}
