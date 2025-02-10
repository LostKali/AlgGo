package main

/*
Given a non-negative integer x, return the square root of x rounded down to the nearest integer.
The returned integer should be non-negative as well.

You must not use any built-in exponent function or operator.

For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.
*/

type Solution struct{}

func (s *Solution) mySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	// 100: 50/50, 25/50, 12/25, 6/12
	bottom := x / 2
	high := bottom

	for bottom*bottom > x {
		pB := bottom
		bottom = high / 2
		high = pB
	}

	// 6/12, 9/12
	for high > bottom && (high-bottom) > 1 {
		c := (high + bottom) / 2
		if c*c > x {
			high = c
		} else {
			bottom = c
		}
	}

	// покурить с точки зрения математического доказательства (сходимость рядов)
	return bottom
}

func main() {
	solution := Solution{}
	res := solution.mySqrt(100)
	print(res)
}
