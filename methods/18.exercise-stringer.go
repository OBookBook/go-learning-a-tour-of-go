/*
Exercise: Stringers
https://go-tour-jp.appspot.com/methods/18

実行コマンド
go run methods/18.exercise-stringer.go
*/

package main

import "fmt"

type IPAddr [4]byte

func (ip IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", ip[0], ip[1], ip[2], ip[3])
}

func main() {
	ip := IPAddr{1, 2, 3, 4}
	fmt.Println(ip.String())
}
