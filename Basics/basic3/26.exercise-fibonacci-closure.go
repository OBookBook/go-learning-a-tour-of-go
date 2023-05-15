/*
Function closures
https://go-tour-jp.appspot.com/moretypes/26

実行コマンド
go run Basics/basic3/26.exercise-fibonacci-closure.go
*/
package main

import "fmt"

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		result := a
		a, b = b, a+b
		return result
	}
}

func main() {
	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(i, fib())
	}
}
