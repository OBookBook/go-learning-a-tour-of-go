/*
Switch with no condition
https://go-tour-jp.appspot.com/flowcontrol/11

実行コマンド
go run Basics/Basic2/11.switch-with-no-condition.go
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// 条件のないswitchは、 switch true と書くことと同じ。
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
