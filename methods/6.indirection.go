/*
Methods and pointer indirection
https://go-tour-jp.appspot.com/methods/6

ポインタレシーバを持つメソッドでは、
変数とポインタの両方を適切に使用することができるため、柔軟な利用が可能です。

実行コマンド
go run methods/6.indirection.go
*/

package main

import "fmt"

type Vertex struct {
	X, Y float64
}

/*
	メソッドがポインタレシーバである場合、
	呼び出し時に、変数、または、ポインタのいずれかのレシーバとして取ることができます

	var v Vertex
	v.Scale(5)  // OK この場合、Go言語は自動的に(&v).Scale(5)と解釈し、ポインタレシーバを使用したメソッドを呼び出します。
	p := &v
	p.Scale(10) // OK
*/
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

// メソッド
func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{3, 4}
	p.Scale(2)
	ScaleFunc(p, 10)

	fmt.Println(v, p) // {60 80} &{60 80}
}
