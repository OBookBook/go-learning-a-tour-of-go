package main

import "fmt"

const (
	//  100 ビット左にシフト
	Big = 1 << 100
	// 99 ビット右にシフト
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println((needInt(Small)))   // 出力結果 21
	fmt.Println((needFloat(Small))) // 出力結果 0.2
	fmt.Println((needFloat(Big)))   // 出力結果 1.2676506002282295e+29
}
