package slidingwindow

/*
Given an array of positive numbers and a positive number 'k,'
find the maximum sum of any contiguous subarray of size 'k'.
*/

type Solution struct{}

func (s *Solution) findMaxSumSubArray(k int, arr []int) int {
	if len(arr) < k {
		panic("Array too small")
	}

	maxSum := 0
	for i := 0; i < k; i++ {
		maxSum += arr[i]
	}

	sum := maxSum
	for i := 1; i <= len(arr)-k; i++ {
		currSum := sum - arr[i-1] + arr[i+k-1]
		sum = currSum
		if sum > maxSum {
			maxSum = sum
		}
	}

	return maxSum
}

func main() {
	solution := Solution{}
	res := solution.findMaxSumSubArray(3, []int{2, 1, 5, 1, 3, 2})
	print(res)

	res = solution.findMaxSumSubArray(1, []int{1, 2, 3, 4, 5})
	print(res)
}
