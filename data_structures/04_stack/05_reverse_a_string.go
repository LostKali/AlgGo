package stack

/*
Given a string, write a function that uses a stack to reverse the string.
The function should return the reversed string.
*/

type Solution struct{}

func (s *Solution) reverseString(sInput string) string {
	input := []rune(sInput)
	size := len(input)
	stack := make([]rune, size)
	for idx, r := range input {
		stack[size-1-idx] = r
	}

	return string(stack)
}
