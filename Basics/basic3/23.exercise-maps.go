/*
Exercise: Maps
https://go-tour-jp.appspot.com/moretypes/23

実行コマンド
go run Basics/basic3/23.exercise-maps.go
*/
package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	counts := make(map[string]int)

	for _, word := range words {
		counts[word]++
	}

	return counts
}

func main() {
	s := "I am learning Go programming language. Go is fun!"
	wordCounts := WordCount(s)

	for word, count := range wordCounts {
		fmt.Printf("%s: %d\n", word, count)
	}
}
