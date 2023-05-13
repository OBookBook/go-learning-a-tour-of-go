/*
Maps
https://go-tour-jp.appspot.com/moretypes/19

Go言語でのマップ（連想配列）
マップのゼロ値は nil

実行コマンド
go run Basics/basic3/19.maps.go
*/

package main

import "fmt"

type Vertex struct {
	// 緯度と経度
	Lat, Long float64
}

// mという名前の変数を定義しています。この変数は、string型のキーとVertex型の値を持つマップです。
var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"]) // 出力結果 {40.68433 -74.39967}
}
