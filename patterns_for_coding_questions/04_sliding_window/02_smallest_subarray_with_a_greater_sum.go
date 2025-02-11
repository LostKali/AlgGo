package slidingwindow

/*
Given an array of positive integers and a number ‘S,’ find the length of
the smallest contiguous subarray whose sum is greater than or equal to 'S'.
Return 0 if no such subarray exists.
*/

type Solution struct{}

func (s *Solution) findMinSubArray(S int, arr []int) int {
	length := 0
	stIdx := 0
	currSum := 0

	for i := 0; i < len(arr); i++ {
		currSum += arr[i]
		for currSum >= S {
			if (i-stIdx) < length || length == 0 {
				length = i - stIdx + 1
			}
			currSum -= arr[stIdx]
			stIdx++
		}
	}

	return length
}
