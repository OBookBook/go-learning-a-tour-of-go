/*
Pointers to structs
https://go-tour-jp.appspot.com/moretypes/4

structのフィールドは、structのポインタを通してアクセスすることもできます。

実行コマンド
go run Basics/basic3/4.struct-pointers.go
*/

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	// 1e9は指数表記を表す数値です。
	// 具体的には、1を10の9乗（10^9）倍した値を表しています。つまり、1e9は1000000000と等価です。
	p.X = 1e9
	fmt.Println(v) // 出力結果 {1000000000 2}
}
