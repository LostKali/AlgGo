package warmup

import (
	"unicode"
)

/*
A phrase is a palindrome if, after converting all uppercase letters into lowercase letters
and removing all non-alphanumeric characters, it reads the same forward and backward.
Alphanumeric characters include letters and numbers.

Given a string s, return true if it is a palindrome, or false otherwise.
*/

type Solution struct{}

func (s *Solution) isPalindrome(str string) bool {
	input := []rune(str)

	onlyLetters := []rune{}
	for _, l := range input {
		if isLetterOrDigit(l) {
			onlyLetters = append(onlyLetters, unicode.ToLower(l))
		}
	}

	left := 0
	right := len(onlyLetters) - 1

	for right > left {
		if onlyLetters[left] != onlyLetters[right] {
			return false
		}
		left++
		right--
	}

	return true // ну вот еще вопрос - надо ли нам пустые строки и строки без разрешенного набора символов считать плалиндромами
}

func isLetterOrDigit(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

func main() {
	solution := Solution{}

	res := solution.isPalindrome("A man, a plan, a canal, Panama!")
	println(res)
}
