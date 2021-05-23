package main

func calculateTriangleMaxSum_noDP(nt *NumbersTriangle, vertexIndex int, calculatedBefore []PossiblyCalculatedValue) *Solution {
	if vertexIndex >= nt.Len {
		return &Solution{
			Value: 0,
			Path:  nil,
		}
	}

	cnt++

	rightMaxSolution := calculateTriangleMaxSum_noDP(nt, getRightVertexIndex(vertexIndex), calculatedBefore)
	leftMaxSolution := calculateTriangleMaxSum_noDP(nt, getLeftVertexIndex(vertexIndex), calculatedBefore)

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

func calculateTriangleMinSum_noDP(nt *NumbersTriangle, vertexIndex int, calculatedBefore []PossiblyCalculatedValue) *Solution {
	if vertexIndex >= nt.Len {
		return &Solution{
			Value: 0,
			Path:  nil,
		}
	}

	cnt++

	rightMinSolution := calculateTriangleMinSum_noDP(nt, getRightVertexIndex(vertexIndex), calculatedBefore)
	leftMinSolution := calculateTriangleMinSum_noDP(nt, getLeftVertexIndex(vertexIndex), calculatedBefore)

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

func (nt *NumbersTriangle) calculateTriangleMinSum_noDP() *Solution {
	return calculateTriangleMinSum_noDP(nt, 0, make([]PossiblyCalculatedValue, nt.Len))
}

func (nt *NumbersTriangle) calculateTriangleMaxSum_noDP() *Solution {
	return calculateTriangleMaxSum_noDP(nt, 0, make([]PossiblyCalculatedValue, nt.Len))
}

func calculateTriangleMinSum_noDPWrap(nt *NumbersTriangle) *Solution {
	return calculateTriangleMinSum_noDP(nt, 0, make([]PossiblyCalculatedValue, nt.Len))
}

func calculateTriangleMaxSum_noDPWrap(nt *NumbersTriangle) *Solution {
	return calculateTriangleMaxSum_noDP(nt, 0, make([]PossiblyCalculatedValue, nt.Len))
}
