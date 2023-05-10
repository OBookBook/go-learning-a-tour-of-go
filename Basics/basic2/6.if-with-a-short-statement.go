/*
If with a short statement
https://go-tour-jp.appspot.com/flowcontrol/5
*/

package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	//  一時変数 v をif文の中に定義することで、変数のスコープが明確になる
	//  math.Pow は x の y 乗
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10), // 出力結果 9
		pow(3, 3, 20), // 出力結果 20
	)
}
