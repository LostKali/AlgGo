package warmup

import "fmt"

/*
Given an array of strings words and two different strings that already exist in
the array word1 and word2, return the shortest distance between these two words in the list.
*/

type Solution struct{}

func (s *Solution) shortestDistance(
	words []string,
	word1 string,
	word2 string,
) int {
	var lstW1Idx int
	var lstW2Idx int
	var distance int
	flags := []bool{false, false, false}
	for idx, w := range words {
		if w == word1 {
			lstW1Idx = idx
			flags[0] = true
			flags[2] = true
		}
		if w == word2 {
			lstW2Idx = idx
			flags[1] = true
			flags[2] = true
		}

		if flags[0] && flags[1] && flags[2] {
			currDist := lstW2Idx - lstW1Idx
			if lstW2Idx < lstW1Idx {
				currDist = currDist * (-1)
			}

			if distance == 0 || currDist < distance {
				distance = currDist
			}
			flags[2] = false
		}
	}

	return distance
}

func main() {
	solution := Solution{}
	distance := solution.shortestDistance(
		[]string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"},
		"fox",
		"dog",
	)
	fmt.Println("Distance: ", distance)
}
