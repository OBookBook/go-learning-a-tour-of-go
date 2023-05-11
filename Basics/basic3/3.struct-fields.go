/*
Struct Fields
https://go-tour-jp.appspot.com/moretypes/3

実行コマンド
go run Basics/basic3/3.struct-fields.go
*/
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	// structのフィールドは、ドット( . )を用いてアクセス
	v.X = 4
	fmt.Println(v.X) // 出力結果 4
}
