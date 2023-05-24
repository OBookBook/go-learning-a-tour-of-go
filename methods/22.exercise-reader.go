/*
Readers
https://go-tour-jp.appspot.com/methods/22

実行コマンド ※ 実行しないでね
go run methods/22.exercise-reader.go
*/
package main

import (
	"fmt"
	"io"
)

type MyReader struct{}

func (r MyReader) Read(b []byte) (int, error) {
	for i := 0; i < len(b); i++ {
		b[i] = 'A'
	}
	return len(b), nil
}

func main() {
	reader := MyReader{}
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err == io.EOF {
			break
		}
		fmt.Printf("Read %d bytes: %q\n", n, buffer[:n])
	}
}
