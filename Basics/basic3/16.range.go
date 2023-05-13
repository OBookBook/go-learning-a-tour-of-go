/*
Range
https://go-tour-jp.appspot.com/moretypes/16

反復処理

実行コマンド
go run Basics/basic3/16.range.go
*/

package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	// i 現在の指数  v 要素 pow 配列
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

/*
	出力結果
		2**0 = 1
		2**1 = 2
		2**2 = 4
		2**3 = 8
		2**4 = 16
		2**5 = 32
		2**6 = 64
		2**7 = 128
*/
