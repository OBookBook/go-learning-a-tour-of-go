/*
Channels
https://go-tour-jp.appspot.com/concurrency/1

処理の流れは次のようになります：

sum関数が2つのゴルーチンとして非同期に実行され、スライスの前半と後半の要素の合計が計算されます。
合計値がチャネルに送信されます。
main関数では、チャネルからの2つの値を受信して、表示します。

実行コマンド
go run concurrency/2.channels.go
*/

package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // v をチャネル ch へ送信する
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	// 2つのゴルーチンを起動します。これにより、スライス s の前半と後半の要素を並行して処理し、結果をチャネル c に送信。
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	// チャネル c から2つの値を受信します。この操作はブロックされ、ゴルーチンからのデータが到着するまで待機
	x, y := <-c, <-c // ch から受信した変数を v へ割り当てる

	fmt.Println(x, y, x+y) // -5 17 12
}
