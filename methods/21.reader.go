/*
Readers
https://go-tour-jp.appspot.com/methods/21

実行コマンド
go run methods/21.reader.go
*/
package main

import (
	"fmt"
	"io"
	"strings"
)

// 与えられた文字列をバイト単位で読み取り、読み取ったバイト数やエラー情報を表示しながら、データの読み取りが完了するまで繰り返す。
func main() {
	r := strings.NewReader("Hello, Reader!")
	// 8バイトの容量を持つバイト配列
	b := make([]byte, 8)
	for {
		// 読み込んだバイト数 n とエラー情報 err
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
