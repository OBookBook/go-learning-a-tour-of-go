/*
Mutating Maps
https://go-tour-jp.appspot.com/moretypes/22

Mapの操作

実行コマンド
go run Basics/basic3/22.mutating-maps.go
*/

package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"]) //出力結果 The value: 42

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"]) //出力結果 The value: 48

	// 要素の削除 delete(m, key)
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"]) //出力結果 The value: 0

	// キーに対する要素が存在するかどうか
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok) //出力結果 The value: 0 Present? false
}
