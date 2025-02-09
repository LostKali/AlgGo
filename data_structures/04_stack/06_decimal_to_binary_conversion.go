package stack

/*
Given a positive integer n, write a function that returns its binary equivalent as a string.
The function should not use any in-built binary conversion function.
*/

type Solution struct{}

func (s *Solution) decimalToBinary(num int) string {
	if num == 0 {
		return "0"
	}

	inverted := []int{}

	for num > 0 {
		inverted = append(inverted, num%2)
		num /= 2
	}

	size := len(inverted)
	res := make([]rune, size)
	for idx, s := range inverted {
		res[size-1-idx] = rune('0' + s)
	}

	return string(res)
}
