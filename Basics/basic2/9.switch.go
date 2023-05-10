/*
Switch
https://go-tour-jp.appspot.com/flowcontrol/9

switch ステートメントは if - else ステートメントのシーケンスを短く書ける。
Go では選択された case だけを実行してそれに続く全ての case は実行さらない。
case の最後に必要な break ステートメントが Go では自動的に提供されます。
*/
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s. \n", os) // 出力結果 windows.
	}
}
