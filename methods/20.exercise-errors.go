/*
Exercise: Errors
https://go-tour-jp.appspot.com/methods/20

実行コマンド
go run methods/20.exercise-errors.go
*/
package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt(9))  // 3 <nil>
	fmt.Println(Sqrt(-2)) // 0 cannot Sqrt negative number: -2
}
