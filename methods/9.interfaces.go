/*
Interfaces
https://go-tour-jp.appspot.com/methods/9

実行コマンド
go run methods/9.interfaces.go
*/
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

// 型 MyFloatはfloat64の型の別名
type MyFloat float64

// MyFloat のメソッド
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	// Abser型の変数を宣言
	var a Abser
	// MyFloat型の値f
	f := MyFloat(-math.Sqrt2)
	// Vertex型の値v
	v := Vertex{3, 4}

	// fはMyFloat型であり、MyFloat型はAbserインターフェースを実装しているため、代入が可能
	a = f
	// &vはVertex型のポインタであり、*Vertex型はAbserインターフェースを実装しているため、代入が可能
	a = &v

	fmt.Println(a.Abs()) // 5
}
