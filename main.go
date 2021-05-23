package main

import (
	"fmt"
	"math/rand"
)

var cnt int

func main() {
	nt := newNumbersTriangle([]int{7, 3, 8, 8, 1, 0, 2, 7, 4, 4, 4, 5, 2, 6, 5})

	nt.print()

	fmt.Println()

	cnt = 0
	calculateTriangleMaxSum(nt, 0, make([]PossiblyCalculatedValue, nt.Len)).print(*nt)
	fmt.Printf("cnt: %d \n", cnt)
	fmt.Println()
	cnt = 0
	calculateTriangleMinSum(nt, 0, make([]PossiblyCalculatedValue, nt.Len)).print(*nt)
	fmt.Printf("cnt: %d \n", cnt)

	fmt.Println()
	fmt.Println()
	fmt.Println()

	var accounts = make(accounts, N)
	for i := 0; i < N; i++ {
		accounts[i] = account{rand.Float64() * GMax}
	}

	fmt.Printf("Greedy algorithm: %s%.2f%s\n", RedColor, task(accounts), EndColor)

}
