/*
For is Go's "while"
https://go-tour-jp.appspot.com/flowcontrol/3
*/
package main

import "fmt"

func main() {
	sum := 1
	// Go言語でwhilは、for だけを使います。
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum) // 出力結果 1024
}
