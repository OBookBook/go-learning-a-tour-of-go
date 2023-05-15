/*
Methods
https://go-tour-jp.appspot.com/methods/1

Goには、クラスの仕組みが無い。
構造体や型に対してメソッドを関連付けることができます。
メソッドは、特定の型に対して実行可能な関数のようなものです。

実行コマンド
go run methods/1.method.go
*/

package main

import (
	"fmt"
	"math"
)

type VerTex struct {
	X, Y float64
}

/*
(v Vertex) がレシーバです。
v はレシーバ変数と呼ばれる、メソッドが呼び出される対象の型のインスタンスを指します。
下記の例では、Vertex 型のインスタンスに対して Abs() メソッドが紐づいています。
*/

func (v VerTex) abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := VerTex{3, 4}
	fmt.Println(v.abs()) // 5
}
