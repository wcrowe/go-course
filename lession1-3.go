//package main //comment out if have more than one main in project

import "fmt"

func square(x int) int {
	return x * x
}

func fibonacci(n int) int {
	a := 0
	b := 1
	for i := 0; i < n-1; i++ {
		a, b = b, a+b // b to a a+b to b
	}
	return b
}

func fibonaccirec(n int) int {
	//base case
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonaccirec(n-1) + fibonaccirec(n-2)
	}

}

func main() {
	fmt.Println(square(2))
	fmt.Println("the n-th fibonacci number is:", fibonacci(5))
	fmt.Println("the n-th fibonacci number is:", fibonaccirec(25))

	// input
	// fmt.Print("Enter your name: ")
	// var name string
	// fmt.Scanln(&name)
	// fmt.Println("Hello", name)

	fmt.Print("Enter your number: ")
	var n int
	fmt.Scanln(&n)
	fmt.Println("The nth fibonacci number is", fibonaccirec(n))
}
