/*
Stacking defers
https://go-tour-jp.appspot.com/flowcontrol/13

「スタック（stack）」は、データの一時的な保存や処理のために使用されるデータ構造の一つ
スタックは「後入れ先出し（Last-In-First-Out; LIFO）」の原則に基づいています。
最後に追加されたデータが最初に削除されるという意味です。


実行コマンド
go run Basics/Basic2/13.defer-multi.go
*/

package main

import "fmt"

func main() {
	fmt.Println(("counting"))

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("done")
}

/*
出力結果
counting
done
9
8
7
6
5
4
3
2
1
0
*/
