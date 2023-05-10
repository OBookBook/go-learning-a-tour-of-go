package main

/*
	BUT
	import "fmt"
	import "math"
*/

// Good!!
import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	// 出力結果 Now you have 2.6457513110645907 problems.
}
