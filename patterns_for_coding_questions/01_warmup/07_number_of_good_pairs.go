package warmup

/*
Given an array of integers nums, return the number of good pairs.

A pair (i, j) is called good if nums[i] == nums[j] and i < j.
*/

type Solution struct{}

func (sol *Solution) numGoodPairs(nums []int) int {
	idxs := make(map[int]int)
	pairCount := 0

	for _, n := range nums {
		idxs[n]++
		// every new occurrence of a number can be paired with every previous occurrence
		// so if a number has already appeared 'p' times, we will have 'p-1' new pairs
		pairCount += idxs[n] - 1
	}

	return pairCount
}
