/*
Range and Close
https://go-tour-jp.appspot.com/concurrency/4

実行コマンド
go run concurrency/4.range-and-close.go
*/
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10) // バッファ付きのチャネル 最大で10個の要素を保持できる
	// cap(c) は、チャネル c の容量（バッファサイズ）を返す組み込み関数
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}

/*
	0
	1
	1
	2
	3
	5
	8
	13
	21
	34
*/
