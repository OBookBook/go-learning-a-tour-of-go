/*
Select
https://go-tour-jp.appspot.com/concurrency/5

select ステートメントは、goroutineを複数の通信操作で待たせます。

実行コマンド
go run concurrency/5.select.go
*/
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		// select ステートメントを使用して複数のチャネルの操作を監視
		select {
		// フィボナッチ数列の数値を送信
		case c <- x:
			// プログラムの終了を通知
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
