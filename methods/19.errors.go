/*
Errors
https://go-tour-jp.appspot.com/methods/19

実行コマンド
go run methods/19.errors.go
*/

package main

import (
	"fmt"
	"time"
)

// MyError型がerrorインターフェースを暗黙的に満たしています。
type MyError struct {
	When time.Time
	What string
}

// errorインターフェースは、Error()というメソッドのみを持っています。
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// MyError型のポインタをerror型として返しています。
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		// fmt.Println(err)を呼び出すと、
		// 内部的にはerrがerrorインターフェースを実装しているかどうかをチェックし、実装している場合にはError()メソッドを呼び出してエラーメッセージを取得します。
		fmt.Println(err) // at 2023-05-21 00:24:27.2928106 +0900 JST m=+0.003035301, it didn't work
	}
}
