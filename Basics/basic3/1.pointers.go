/*
Pointers
https://go-tour-jp.appspot.com/moretypes/1

ポインタとは、値のメモリアドレスを指す
&変数名 変数のメモリアドレス を取得
*変数名 値を取り出す。

実行コマンド
go run Basics/basic3/1.pointers.go
*/

package main

import "fmt"

func main() {
	i, j := 42, 2701

	// 変数 i のポインタを変数 p に代入
	p := &i
	// *オペレータは、ポインタが指しているメモリ上の変数を示す。
	// ポインタ p を通じて変数 i の値を表示
	fmt.Println(*p) // 42
	// ポインタ p を通じて変数 i に値 21 を代入
	*p = 21
	fmt.Println(i) // 21

	p = &j
	*p = *p / 37
	fmt.Println(j) // 73
}
