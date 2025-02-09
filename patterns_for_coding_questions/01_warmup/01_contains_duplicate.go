package warmup

/*
Given an integer array nums, return true if any value appears at least twice in the array,
and return false if every element is distinct.
*/

type Solution struct{}

func (s *Solution) containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	for _, n := range nums {
		if _, exists := m[n]; exists {
			return true
		}
		m[n] = n
	}

	return false
}
