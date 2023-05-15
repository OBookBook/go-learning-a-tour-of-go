/*
Function values
https://go-tour-jp.appspot.com/moretypes/24

コールバック関数

実行コマンド
go run Basics/basic3/24.function-values.go
*/

package main

import (
	"fmt"
	"math"
)

// 関数を引数に取る。
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}

	fmt.Println(hypot(5, 12))      //出力結果 13
	fmt.Println(compute(hypot))    //出力結果 5
	fmt.Println(compute(math.Pow)) //出力結果 81
}
