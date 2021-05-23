package lab9

import (
	"fmt"
	"math"
)

type NumbersTriangle struct {
	Array []int
	Len   int
}

type PossiblyCalculatedValue struct {
	Calculated bool
	Value      *Solution
}

type Solution struct {
	Value int
	Path  []int
}

//returns height of the triangle represented by arr
/*
1
2 3
4 5 6
-> 3
*/
func getHeight(arr []int) int {
	return int(math.Floor(math.Sqrt(float64(2*len(arr))+0.25) - 0.5))
}

func checkIfTriangle(arr []int) {
	n := getHeight(arr)
	if n*(n+1)/2 != len(arr) {
		panic("not a valid triangle (length is not equal to height*(height+1)/2 (see arithmetic progression sum)")
	}
}

func newNumbersTriangle(arr []int) *NumbersTriangle {
	checkIfTriangle(arr)
	return &NumbersTriangle{Array: arr, Len: len(arr)}
}

func (nt *NumbersTriangle) getHeight() int {
	return int(math.Sqrt((float64(2*len(nt.Array)) + 0.25) - 0.5))
}

func (nt *NumbersTriangle) printTriangle() {
	for i := 0; i < nt.getHeight(); i++ {
		for j := 0; j < i+1; j++ {
			fmt.Printf("%d ", nt.get(j+i*(1+i)/2))
		}
		fmt.Println()

	}
}

func getLayer(index int) int {
	sum := 0
	var layer = 1
	for index >= sum {
		sum += layer
		layer++
	}
	return layer - 1
}

func getLeftVertexIndex(index int) int {
	return index + getLayer(index)
}

func getRightVertexIndex(index int) int {
	return getLeftVertexIndex(index) + 1
}

func (nt *NumbersTriangle) get(index int) int {
	return nt.Array[index]
}

func (sol *Solution) print(nt NumbersTriangle) {
	fmt.Println(sol.Value)
	for i := 0; i < len(sol.Path); i++ {
		for j := 0; j < i+1; j++ {
			if j+i*(1+i)/2 == sol.Path[len(sol.Path)-i-1] {
				fmt.Printf("%d ", nt.get(j+i*(1+i)/2))
			} else {
				fmt.Print("* ")
			}

		}
		fmt.Println()

	}
}

func calculateTriangleMaxSum(nt *NumbersTriangle, vertexIndex int, calculatedBefore []PossiblyCalculatedValue) *Solution {
	if vertexIndex >= nt.Len {
		return &Solution{
			Value: 0,
			Path:  nil,
		}
	}

	//DP part
	if calculatedBefore[vertexIndex].Calculated {
		return calculatedBefore[vertexIndex].Value
	}
	//

	cnt++

	rightMaxSolution := calculateTriangleMaxSum(nt, getRightVertexIndex(vertexIndex), calculatedBefore)
	leftMaxSolution := calculateTriangleMaxSum(nt, getLeftVertexIndex(vertexIndex), calculatedBefore)

	sum := nt.get(vertexIndex)
	var path []int
	if rightMaxSolution.Value >= leftMaxSolution.Value {
		sum += rightMaxSolution.Value
		path = append(rightMaxSolution.Path, vertexIndex)
	} else {
		sum += leftMaxSolution.Value
		path = append(leftMaxSolution.Path, vertexIndex)
	}

	calculatedBefore[vertexIndex].Calculated = true
	calculatedBefore[vertexIndex].Value = &Solution{
		Value: sum,
		Path:  path,
	}
	return calculatedBefore[vertexIndex].Value
}

func calculateTriangleMinSum(nt *NumbersTriangle, vertexIndex int, calculatedBefore []PossiblyCalculatedValue) *Solution {
	if vertexIndex >= nt.Len {
		return &Solution{
			Value: 0,
			Path:  nil,
		}
	}

	//DP part
	if calculatedBefore[vertexIndex].Calculated {
		return calculatedBefore[vertexIndex].Value
	}
	//

	cnt++

	rightMinSolution := calculateTriangleMinSum(nt, getRightVertexIndex(vertexIndex), calculatedBefore)
	leftMinSolution := calculateTriangleMinSum(nt, getLeftVertexIndex(vertexIndex), calculatedBefore)

	sum := nt.get(vertexIndex)
	var path []int
	if rightMinSolution.Value <= leftMinSolution.Value {
		sum += rightMinSolution.Value
		path = append(rightMinSolution.Path, vertexIndex)
	} else {
		sum += leftMinSolution.Value
		path = append(leftMinSolution.Path, vertexIndex)
	}

	calculatedBefore[vertexIndex].Calculated = true
	calculatedBefore[vertexIndex].Value = &Solution{
		Value: sum,
		Path:  path,
	}
	return calculatedBefore[vertexIndex].Value
}
