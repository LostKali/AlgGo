package main

/*
Given a square matrix (2D array), calculate the sum of its two diagonals.

The two diagonals in consideration are the primary diagonal that spans from the top-left
to the bottom-right and the secondary diagonal that spans from top-right to bottom-left.
If a number is part of both diagonals (which occurs only for odd-sized matrices),
it should be counted only once in the sum.

Constraints:

n == mat.length == mat[i].length
1 <= n <= 100
1 <= mat[i][j] <= 100
*/

import "fmt"

type Solution struct{}

func (s Solution) diagonalSum(mat [][]int) int {
	totalSum := 0

	for i := 0; i < len(mat); i++ {
		totalSum += mat[i][i]
		if i != len(mat)-1-i {
			totalSum += mat[i][len(mat)-1-i]
		}
	}

	return totalSum
}

func main() {
	s := Solution{}
	fmt.Println(s.diagonalSum([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}})) // Output: 25
	fmt.Println(s.diagonalSum([][]int{{1, 0}, {0, 1}}))                  // Output: 2
	fmt.Println(s.diagonalSum([][]int{{5}}))                             // Output: 5
}
