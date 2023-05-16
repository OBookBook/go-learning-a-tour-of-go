/*
Pointer receivers
https://go-tour-jp.appspot.com/methods/4

ポインタレシーバでメソッドを宣言できます。

実行コマンド
go run methods/4.methods-pointers.go
*/

package main

import (
	"fmt"
	"math"
)

type VerTex struct {
	X, Y float64
}

func (v VerTex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y) // 50
}

func (v *VerTex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := VerTex{3, 4}
	// ポインタレシーバでメソッドを宣言できます。
	v.Scale(10)
	fmt.Println(v.Abs()) // 50
}
