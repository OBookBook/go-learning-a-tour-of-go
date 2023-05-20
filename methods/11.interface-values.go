/*
Interface values
https://go-tour-jp.appspot.com/methods/11

実行コマンド
go run methods/11.interface-values.go
*/
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i) // (&{Hello}, *main.T)
	i.M()       // Hello

	i = F(math.Pi)
	describe(i) // (3.141592653589793, main.F)
	i.M()       // 3.141592653589793
}
