/*
Map literals continued
https://go-tour-jp.appspot.com/moretypes/21

もし、mapに渡すトップレベルの型が単純な型名である場合は、
リテラルの要素から推定できますので、その型名を省略することができます。

実行コマンド
go run Basics/basic3/21.map-literals-continued.go
*/

package main

import "fmt"

type Vertex struct {
	Let, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m) // 出力結果 map[Bell Labs:{40.68433 -74.39967} Google:{37.42202 -122.08408}]
}
