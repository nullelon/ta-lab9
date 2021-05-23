package main

import (
	"fmt"
)

func main() {
	nt := newNumbersTriangle([]int{7, 3, 8, 8, 1, 0, 2, 7, 4, 4, 4, 5, 2, 6, 5})

	nt.printTriangle()

	fmt.Println()

	cnt = 0
	calculateTriangleMaxSum(nt, 0, make([]PossiblyCalculatedValue, nt.Len)).print(*nt)
	fmt.Printf("cnt: %d \n", cnt)
	fmt.Println()
	cnt = 0
	calculateTriangleMinSum(nt, 0, make([]PossiblyCalculatedValue, nt.Len)).print(*nt)
	fmt.Printf("cnt: %d \n", cnt)
}
