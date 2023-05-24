/*
Images
https://go-tour-jp.appspot.com/methods/24

実行コマンド
go run methods/24.images.go
*/
package main

import (
	"fmt"
	"image"
)

func main() {
	// 100x100ピクセルのRGBA画像を作成
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	// 作成した画像の範囲を示す矩形
	fmt.Println(m.Bounds()) // (0,0)-(100,100)
	// 画像の座標(0, 0)のピクセルのRGBA値を取得
	fmt.Println(m.At(0, 0).RGBA()) // 0 0 0 0
}
