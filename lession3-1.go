//package main

import (
	"fmt"
	"strconv"
)

func euclid(a int, b int) int {
	for a != b {
		if a < b {
			b = b - a
		} else {
			a = a - b
		}
	}
	return a
}

type fraction struct {
	numerator   int
	denominator int
}

func (a fraction) add(b fraction) fraction {
	sumOfCrossProducts := a.numerator*b.denominator + b.numerator*a.denominator
	productOfDenominators := a.denominator * b.denominator
	gcdOfNumeratorAndDenominator := euclid(sumOfCrossProducts, productOfDenominators)
	return fraction{sumOfCrossProducts / gcdOfNumeratorAndDenominator, productOfDenominators / gcdOfNumeratorAndDenominator}
}

func (a fraction) subtract(b fraction) fraction {
	differenceOfCrossProducts := a.numerator*b.denominator - b.numerator*a.denominator
	productOfDenominators := a.denominator * b.denominator
	gcdOfNumeratorAndDenominator := euclid(differenceOfCrossProducts, productOfDenominators)
	return fraction{differenceOfCrossProducts / gcdOfNumeratorAndDenominator, productOfDenominators / gcdOfNumeratorAndDenominator}
}

func (a fraction) multiply(b fraction) fraction {
	productOfNumerators := a.numerator * b.numerator
	productOfDenominators := a.denominator * b.denominator
	gcdOfNumeratorAndDenominator := euclid(productOfNumerators, productOfDenominators)
	return fraction{productOfNumerators / gcdOfNumeratorAndDenominator, productOfDenominators / gcdOfNumeratorAndDenominator}
}

func (a fraction) divide(b fraction) fraction {
	productOfNumerators := a.numerator * b.denominator
	productOfDenominators := a.denominator * b.numerator
	gcdOfNumeratorAndDenominator := euclid(productOfNumerators, productOfDenominators)
	return fraction{productOfNumerators / gcdOfNumeratorAndDenominator, productOfDenominators / gcdOfNumeratorAndDenominator}
}

func (a fraction) toString() string {
	return strconv.Itoa(a.numerator) + "/" + strconv.Itoa(a.denominator)
}

func main() {
	// This is the main function
	// this is a tes
	oneHalf := fraction{1, 2}
	oneThird := fraction{1, 3}
	// one
	fmt.Println("1/2 + 1/3 = ", oneHalf.add(oneThird).toString())
	fmt.Println("1/2 - 1/3 = ", oneHalf.subtract(oneThird).toString())
	fmt.Println("1/2 * 1/3 = ", oneHalf.multiply(oneThird).toString())
	fmt.Println("1/2 / 1/3 = ", oneHalf.divide(oneThird).toString())

	fiveSixth := fraction{5, 6}
	sevenEights := fraction{7, 8}
	// one
	fmt.Println("5/6 + 7/8 = ", fiveSixth.add(sevenEights).toString())
	fmt.Println("5/6 - 7/8 = ", fiveSixth.subtract(sevenEights).toString())
	fmt.Println("5/6  * 7/8 = ", fiveSixth.multiply(sevenEights).toString())
	fmt.Println("5/6 / 7/8 = ", fiveSixth.divide(sevenEights).toString())
}
