package main

/*
Given an input array of integers nums, find an integer array, let's call it differenceArray,
of the same length as an input integer array.

Each element of differenceArray, i.e., differenceArray[i], should be calculated as follows:
take the sum of all elements to the left of index i in array nums (let's call it leftSumi),
 and subtract it from the sum of all elements to the right of index i in array nums
 (let's call it rightSumi), taking the absolute value of the result:

differenceArray[i] = | leftSum_i - rightSum_i |

If there are no elements to the left or right of i, the corresponding sum should be taken as 0.
*/

type Solution struct{}

func (s *Solution) findDifferenceArray(nums []int) []int {
	if nums == nil || len(nums) == 0 {
		return []int{}
	}

	n := len(nums)
	differenceArray := make([]int, n)

	var leftSum, rightSum int = 0, 0

	for _, n := range nums {
		rightSum += n
	}

	for i, n := range nums {
		rightSum = rightSum - leftSum - n
		differenceArray[i] = abs(leftSum - rightSum)
		leftSum += n
	}

	return differenceArray
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
