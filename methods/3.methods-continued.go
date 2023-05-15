/*
Methods continued
https://go-tour-jp.appspot.com/methods/3

任意の型にもメソッドを宣言することができます。
具体的な例として、数値型の MyFloat 型に Abs メソッドを追加する方法を示しています。

実行コマンド
go run methods/3.methods-continued.go
*/

package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f)       // -1.4142135623730951
	fmt.Println(f.abs()) // 1.4142135623730951
}
