/*
If
https://go-tour-jp.appspot.com/flowcontrol/5
*/

package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	// マイナスの値がきたら、再帰呼び出しが行われている。
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4)) // 出力結果 1.4142135623730951 2i
}
