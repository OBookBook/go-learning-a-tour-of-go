/*
Creating a slice with make
https://go-tour-jp.appspot.com/moretypes/13

cap（容量）は、スライスが参照している元の配列内で保持できる要素の最大数を表します。
スライスの容量は、元の配列の先頭からの要素数となります。

実行コマンド
go run Basics/basic3/13.making-slices.go
*/

package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a) // 出力結果 a len=5 cap=5 [0 0 0 0 0]

	b := make([]int, 0, 5)
	printSlice("b", b) // 出力結果 b len=0 cap=5 []

	// スライスbから最初から2番目の要素までを切り出してスライスcを作成します
	c := b[:2]
	printSlice("c", c) // 出力結果 c len=2 cap=5 [0 0]

	// スライスcから2番目から5番目の要素までを切り出してスライスdを作成します
	// cap(d) = cap(c) - start
	d := c[2:5]
	printSlice("d", d) // 出力結果 d len=3 cap=3 [0 0 0]
}

/*
	% フォーマット指定子
		%s：文字列を表示するための指定子。sにはs型の引数が入ります。
		%d：整数を表示するための指定子。dには整数型の引数が入ります。
		%v：任意の型の値を表示するための指定子。vには値の型に応じた引数が入ります。
*/
func printSlice(s string, x []int) {
	// スライスの長さ(len)、容量(cap)、および要素を出力します
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
