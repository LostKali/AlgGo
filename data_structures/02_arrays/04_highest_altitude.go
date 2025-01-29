package main

/*
A bike rider is going on a ride. The road contains n + 1 points at different altitudes.
The rider starts from point 0 at an altitude of 0.

Given an array of integers gain of length n, where gain[i] represents the net gain in
altitude between points i and i + 1 for all (0 <= i < n), return the highest altitude of a point.
*/
type Solution struct{}

func (s *Solution) largestAltitude(gain []int) int {
	if gain == nil || len(gain) == 0 {
		return 0
	}

	maxAltitude := 0
	altitude := 0

	for _, g := range gain {
		altitude += g
		if altitude > maxAltitude {
			maxAltitude = altitude
		}
	}

	return maxAltitude
}
