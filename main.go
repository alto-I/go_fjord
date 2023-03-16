package main

import (
	"fmt"
	"go_fjord/bowling"
	"os"
	"strings"
)

func main() {
	// fizzbuzz
	// for i := 1; i < 101; i++ {
	// 	fmt.Println(fizzbuzz.Fizzbuzz(i))
	// }

	// ボウリングスコア計算
	slice := strings.Split(os.Args[1], ",")
	fmt.Println(bowling.CalcScore(slice))
}
