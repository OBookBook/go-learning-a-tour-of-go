/*
Slices of slices
https://go-tour-jp.appspot.com/moretypes/14

実行コマンド
go run Basics/basic3/14.slices-of-slice.go
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		// strings.Join() 配列要素をスペースで区切って1つの文字列に結合
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}

	/*
		出力結果
			X _ X
			O _ X
			_ _ O
	*/
}
