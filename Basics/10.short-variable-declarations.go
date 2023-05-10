package main

import "fmt"

func main() {
	// 型 宣言
	var i, j int = 1, 2
	// 暗黙的な型宣言 ※ 関数内でのみ可能 :=
	k := 3
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java) // 出力結果 1 2 3 true false no!
}
