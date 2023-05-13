/*
Map literals
https://go-tour-jp.appspot.com/moretypes/20

mapリテラルは、structリテラルに似ていますが、 キー ( key )が必要です。

実行コマンド
go run Basics/basic3/20.map-literals.go
*/

package main

import "fmt"

type Vertex struct {
	Let, Long float64
}

// Vertex{} 内で Vertex の型が不要な為 修正
var m = map[string]Vertex{
	// "Bell Labs": Vertex{
	"Bell Labs": {
		40.68433, -74.39967,
	},
	// "Google": Vertex{
	"Google": {
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m) // 出力結果 map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
}
