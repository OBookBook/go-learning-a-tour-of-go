/*
If and else
https://go-tour-jp.appspot.com/flowcontrol/7
*/

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		return lim
	}
}

func main() {
	fmt.Println(
		pow(3, 2, 10), // 出力結果 9
		pow(3, 3, 20), // 出力結果 20
	)
}
