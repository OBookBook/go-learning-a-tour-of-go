/*
	Type conversions 型変換
	https://go-tour-jp.appspot.com/basics/13
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z) // 出力結果 3 4 5
}
