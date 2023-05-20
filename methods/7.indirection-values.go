/*
Methods and pointer indirection (2)
https://go-tour-jp.appspot.com/methods/7

実行コマンド
go run methods/7.indirection-values.go
*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())    // 5
	fmt.Println(AbsFunc(v)) // 5

	p := &Vertex{3, 4}
	fmt.Println(p.Abs()) // 5
	// *p を使用してポインタをデリファレンスし、Vertex の値を取得してから AbsFunc に渡しています。
	fmt.Println(AbsFunc(*p)) // 5
}
