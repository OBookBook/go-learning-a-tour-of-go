/*
Struct Literals
https://go-tour-jp.appspot.com/moretypes/5

フィールドの一部だけを列挙することができます

実行コマンド
go run Basics/basic3/5.struct-literals.go
*/

package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	// Xフィールドのみ初期化
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, v2, v3, p)
	// 出力結果{1 2} {1 0} {0 0} &{1 2}
}
