/*
	Zero values
	https://go-tour-jp.appspot.com/basics/12
*/

package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	fmt.Printf("%v %v %v %q\n", i, f, b, s) // 出力結果 0 0 false ""
}
