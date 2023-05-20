/*
Choosing a value or pointer receiver
https://go-tour-jp.appspot.com/methods/8

ポインタレシーバを使う2つの理由
1.メソッドがレシーバが指す先の変数を変更するため：
ポインタレシーバを使用すると、メソッド内での変更がレシーバが指す実際の変数に反映されます。つまり、メソッドがオブジェクトの状態を変更することができます。

2.メソッドの呼び出し毎に変数のコピーを避けるため：
値レシーバを使用すると、メソッド呼び出しの際にレシーバのコピーが作成されます。しかし、ポインタレシーバを使用すると、メソッドは変数のポインタを受け取るため、大きな構造体などの場合に効率的です。ポインタを介して直接データにアクセスできるため、コピーのオーバーヘッドを避けることができます。

実行コマンド
go run methods/8.methods-with-pointer-receivers.go
*/

package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)

}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
