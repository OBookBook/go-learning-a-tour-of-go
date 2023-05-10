/*
https://go-tour-jp.appspot.com/basics/11
*/
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)     // 出力結果 Type: bool Value: false
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt) // 出力結果 Type: uint64 Value: 18446744073709551615
	fmt.Printf("Type: %T Value: %v\n", z, z)           // 出力結果 Type: complex128 Value: (2+3i)
}

/*
Go言語の基本型(組み込み型):

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 の別名

rune // int32 の別名
     // Unicode のコードポイントを表す

float32 float64

complex64 complex128

bool: 真偽値を表す型で、trueまたはfalseの値を取ります。
string: 文字列を表す型で、複数の文字からなる文字列を表現することができます。
int, int8, int16, int32, int64: 整数を表す型で、それぞれ8ビット、16ビット、32ビット、64ビットの符号付き整数を表現することができます。intはプラットフォームによってサイズが異なります。
uint, uint8, uint16, uint32, uint64, uintptr: 整数を表す型で、それぞれ8ビット、16ビット、32ビット、64ビットの符号なし整数を表現することができます。uintptrはメモリアドレスを表します。
byte: uint8の別名です。主にバイト列を表現する場合に使用されます。
rune: int32の別名で、Unicodeのコードポイントを表現するために使用されます。
float32, float64: 浮動小数点数を表す型で、それぞれ32ビット、64ビットの浮動小数点数を表現することができます。
complex64, complex128: 複素数を表す型で、それぞれ実数部と虚数部が32ビット、64ビットの複素数を表現することができます。
*/
