/*
Range continued
https://go-tour-jp.appspot.com/moretypes/17

インデックスや値は、 " _ "(アンダーバー) へ代入することで捨てることができます。
for i, _ := range pow
for _, value := range pow

実行コマンド
go run Basics/basic3/17.range-continued.go
*/
package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
