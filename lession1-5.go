//package main

import (
	"fmt"
	"math"
)

func squareRoot(area float64) float64 {
	guess := 1.0
	otherside := area / guess

	const THRESHOLD = 0

	for math.Abs(guess-otherside) > THRESHOLD {
		guess = (otherside + guess) / 2
		otherside = area / guess
	}
	return guess
}

// same but recursive
func squareRootRec(area float64, guess float64) float64 {

	otherside := area / guess
	threshold := .00001
	if math.Abs(otherside-guess) < threshold {
		return guess
	} else {
		newGuess := (otherside + guess) / 2
		return squareRootRec(area, newGuess)
	}

}

func main() {
	fmt.Print("Enter a number to get sq root: ")
	var x float64
	fmt.Scanln(&x)
	println(squarRootRec(x, 1.0))

}
