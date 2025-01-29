package main

import (
	"fmt"
	"sort"
)

/*
Given an integer array nums, return true if any value appears at least twice in the array,
and return false if every element is distinct.
*/

type Solution struct{}

// containsDuplicate checks for duplicates in a slice of integers
func (s *Solution) containsDuplicate(nums []int) bool {
	if nums == nil || len(nums) == 0 || len(nums) == 1 {
		return false
	}

	sort.Ints(nums)
	var previous = nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == previous {
			return true
		}
		previous = nums[i]
	}
	return false
}

func main() {
	solution := Solution{}

	nums1 := []int{1, 2, 3, 4}
	fmt.Println(solution.containsDuplicate(nums1)) // Expected output: false

	nums2 := []int{1, 2, 3, 1}
	fmt.Println(solution.containsDuplicate(nums2)) // Expected output: true

	nums3 := []int{}
	fmt.Println(solution.containsDuplicate(nums3)) // Expected output: false

	nums4 := []int{1, 1, 1, 1}
	fmt.Println(solution.containsDuplicate(nums4)) // Expected output: true
}
