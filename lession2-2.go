//package main

import "fmt"

func fibonacciMemo(n int) int {
	s := []int{0, 1}
	for i := 0; i < n-1; i++ {
		s = append(s, s[i]+s[i+1])
	}
	return s[len(s)-1]

}

func main() {
	s := make([]int, 2)
	fmt.Println(s)

	s[0] = 1
	s[1] = 2
	fmt.Println(s)

	t := []string{"one", "two"}
	fmt.Println(t)
	fmt.Println(fibonacciMemo(25))
}
