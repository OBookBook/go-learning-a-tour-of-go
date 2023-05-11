/*
Structs
https://go-tour-jp.appspot.com/moretypes/2

struct (構造体)は、フィールド( field )の集まり。

実行コマンド
go run Basics/basic3/2.structs.go
*/
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	// Vertexのインスタンスを作成し、初期値を指定
	v := Vertex{1, 2}
	// vを表示
	fmt.Println(v) // 出力結果 {1 2}
}
