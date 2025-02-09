package warmup

/*
Given a string s, reverse only all the vowels in the string and return it.

The vowels are 'a', 'e', 'i', 'o', and 'u', and they can appear in both lower and upper cases,
more than once.
*/

type Solution struct{}

func (s *Solution) reverseVowels(str string) string {
	stack := []rune{}

	input := []rune(str)

	vowels := make(map[rune]rune)
	vowels['a'] = 'a'
	vowels['A'] = 'A'
	vowels['e'] = 'e'
	vowels['E'] = 'E'
	vowels['i'] = 'i'
	vowels['I'] = 'I'
	vowels['I'] = 'I'
	vowels['o'] = 'o'
	vowels['O'] = 'O'
	vowels['u'] = 'u'
	vowels['U'] = 'U'

	for _, l := range input {
		if _, exists := vowels[l]; exists {
			stack = append(stack, l)
		}
	}

	top := len(stack) - 1
	for idx, l := range input {
		if _, exists := vowels[l]; exists {
			input[idx] = stack[top]
			top--
		}
	}

	return string(input)

}

// todo: возможно решение в один проход с two-pointers
