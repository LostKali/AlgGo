package arrays

import "fmt"

/*
Given a one-dimensional array of integers, create a new array that represents
the running sum of the original array.

The running sum at position i in the new array is calculated as the sum of all
the numbers in the original array from the 0th index up to the i-th index (inclusive).
Formally, the resulting array should be computed as follows:
    result[i] = sum(nums[0] + nums[1] + ... + nums[i])
    for each i from 0 to the length of the array minus one.
*/

type Solution struct{}

func (s *Solution) runningSum(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return []int{}
	}

	result := make([]int, len(nums))
	for i, x := range nums {
		if i == 0 {
			result[i] = x
		} else {
			result[i] = x + result[i-1]
		}
	}
	return result
}

func mainx() {
	solution := &Solution{}

	// Test cases
	testInputs := [][]int{
		{1, 2, 3, 4},
		{3, 1, 4, 2, 2},
		{-1, -2, -3, -4, -5},
	}

	for _, input := range testInputs {
		output := solution.runningSum(input)

		// Print the output array
		for _, val := range output {
			fmt.Print(val, " ")
		}
		fmt.Println()
	}
}
