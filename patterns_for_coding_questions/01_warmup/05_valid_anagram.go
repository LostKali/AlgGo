package warmup

/*
Given two strings s and t, return true if t is an anagram of s, and false otherwise.

An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
typically using all the original letters exactly once.
*/

type Solution struct{}

func (sol *Solution) isAnagram(s, t string) bool {
	sIn := []rune(s)
	tIn := []rune(t)

	letters := make(map[rune]int)
	for _, r := range sIn {
		if _, exists := letters[r]; exists {
			letters[r] = letters[r] + 1
		} else {
			letters[r] = 1
		}
	}

	for _, r := range tIn {
		if _, exists := letters[r]; !exists {
			return false
		}

		c := letters[r] - 1
		if c == 0 {
			delete(letters, r)
		} else {
			letters[r] = c
		}
	}

	return len(letters) == 0
}
