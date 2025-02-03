package main

/*
Given a string s containing (, ), [, ], {, and } characters. Determine if a given string
of parentheses is balanced.

A string of parentheses is considered balanced if every opening parenthesis has a corresponding
closing parenthesis in the correct order.
*/

import (
	"fmt"
	"log"
)

type Solution struct{}

func (s *Solution) isValid(s1 string) bool {
	m := make(map[rune]rune)
	m[')'] = '('
	m[']'] = '['
	m['}'] = '{'

	stack := []rune{}

	for _, r := range s1 {
		if op, exists := m[r]; exists {
			log.Print("Checking: ", string(r))

			if len(stack) == 0 {
				return false
			}

			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if last != op {
				return false
			}
		} else {
			log.Println("Adding to the stack: ", string(r))
			stack = append(stack, r)
		}
	}

	return len(stack) == 0
}

func main() {
	s := Solution{}

	test1 := "{[()]}"
	test2 := "{[}]"
	test3 := "(]"

	fmt.Println("Test 1:", s.isValid(test1))
	fmt.Println("Test 2:", s.isValid(test2))
	fmt.Println("Test 3:", s.isValid(test3))
}
