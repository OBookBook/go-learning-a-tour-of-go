/*
Goroutines
https://go-tour-jp.appspot.com/concurrency/1

平行処理

実行コマンド
go run concurrency/1.goroutines.go
*/
package main

import (
	"fmt"
	"time"
)

// 各出力の間には100ミリ秒の待機
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("world") // 非同期処理
	say("hello")    // メインゴルーチン
}
