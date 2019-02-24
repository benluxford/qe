package qubit

import "math"

// Max : Return the max probability
func Max(input []float64) (max float64) {
	// set max to the first value
	max = input[0]
	// for all numbers in input slice
	for _, number := range input {
		// return the max of the two numbers and set max to highest
		max = math.Max(max, number)
	}
	// return the highest number
	return
}

// Min : Return the min probability
func Min(input []float64) (min float64) {
	// set min to the first value
	min = input[0]
	// for all numbers in input slice
	for _, number := range input {
		// return the max of the two numbers and set max to highest
		min = math.Min(min, number)
	}
	// return the lowest number
	return min
}

// func Sum(p []float64) float64 {
// 	var sum float64
// 	for _, pp := range p {
// 		sum = sum + pp
// 	}

// 	return sum
// }
