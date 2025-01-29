package main

/*
Given a binary matrix that has dimensions , consisting of ones and zeros,
determine the row that contains the highest number of ones and return two values:
the zero-based index of this row and the actual count of ones it possesses.

If there is a tie, i.e., multiple rows contain the same maximum number of ones,
we must select the row with the lowest index.
*/

import "fmt"

type Solution struct{}

func (s *Solution) rowAndMaximumOnes(mat [][]int) []int {
	maxOnesCount := 0
	maxOnesRowIdx := 0

	for idx, row := range mat {
		onesCount := 0
		for _, el := range row {
			if el == 1 {
				onesCount += el
			}
		}

		if onesCount > maxOnesCount {
			maxOnesCount = onesCount
			maxOnesRowIdx = idx
		}
	}

	return []int{maxOnesRowIdx, maxOnesCount}
}

func main() {
	sol := Solution{}
	// Applying example inputs
	result1 := sol.findMaxOnesRow([][]int{{1, 0}, {1, 1}, {0, 1}})
	fmt.Println(result1[0], ",", result1[1]) // Output: 1, 2

	result2 := sol.findMaxOnesRow([][]int{{0, 1, 1}, {0, 1, 1}, {1, 1, 1}})
	fmt.Println(result2[0], ",", result2[1]) // Output: 2, 3

	result3 := sol.findMaxOnesRow([][]int{{1, 0, 1}, {0, 0, 1}, {1, 1, 0}})
	fmt.Println(result3[0], ",", result3[1]) // Output: 0, 2
}