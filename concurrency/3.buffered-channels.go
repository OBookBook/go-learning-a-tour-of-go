/*
Buffered Channels
https://go-tour-jp.appspot.com/concurrency/3

チャネルは、 バッファ ( buffer )として使えます。
バッファが詰まった時は、チャネルへの送信をブロックします。
バッファが空の時には、チャネルの受信をブロックします。
実行コマンド
go run concurrency/3.buffered-channels.go
*/

package main

import "fmt"

func main() {
	// バッファを持つチャネルを初期化するには、 make の２つ目の引数にバッファの長さを与えます:
	ch := make(chan int, 2) // バッファ付きチャネル
	ch <- 1                 // バッファに要素を送信
	ch <- 2                 // バッファに要素を送信
	fmt.Println(<-ch)       // バッファから要素を受信  1
	fmt.Println(<-ch)       // バッファから要素を受信  2
}
