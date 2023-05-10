/*
For continued
https://go-tour-jp.appspot.com/flowcontrol/2
*/
package main

import "fmt"

func main() {
	sum := 1
	// 初期化と後処理ステートメントの記述は任意
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum) // 出力結果 1024
}
