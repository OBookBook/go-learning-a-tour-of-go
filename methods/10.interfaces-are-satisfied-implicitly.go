/*
Interfaces are implemented implicitly
https://go-tour-jp.appspot.com/methods/10

実行コマンド
go run methods/10.interfaces-are-satisfied-implicitly.go
*/

package main

import "fmt"

// インターフェース
type I interface {
	M()
}

// 構造体
type T struct {
	S string
}

// このメソッドは、T型がインターフェースIを実装していることを意味します、
// しかし、そのことを明示的に宣言する必要はない。
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M() // hello

	var you T = T{"world"}
	you.M()
}
