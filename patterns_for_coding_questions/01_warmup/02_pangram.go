package main

/*
A pangram is a sentence where every letter of the English alphabet appears at least once.

Given a string sentence containing English letters (lower or upper-case), return true
if sentence is a pangram, or false otherwise.

Note: The given sentence might contain other characters like digits or spaces, your solution
should handle these too.
*/

type Solution struct{}

func (s *Solution) checkIfPangram(sentence string) bool {
	alphabet := make(map[rune]rune)
	for r := 'A'; r <= 'Z'; r++ {
		alphabet[r] = r
	}

	input := []rune(sentence)

	for _, l := range input {
		var lUp rune
		if 'a' <= l && l <= 'z' {
			lUp = l - ('a' - 'A')
		} else {
			lUp = l
		}
		if _, exists := alphabet[lUp]; exists {
			delete(alphabet, lUp)
		}
	}

	return len(alphabet) == 0
}

func main() {
	solution := Solution{}

	res1 := solution.checkIfPangram("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	println(res1)

	res2 := solution.checkIfPangram("This is not a pangram")
	println(res2)

	res3 := solution.checkIfPangram("TheQuickBrownFoxJumpsOverTheLazyDog")
	println(res3)
}
