/*
Methods are functions
https://go-tour-jp.appspot.com/methods/2

実行コマンド
go run methods/2.methods-funcs.go
*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// この Abs は、先ほどの例から機能を変えずに通常の関数として記述しています。
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v)) // 5
}
