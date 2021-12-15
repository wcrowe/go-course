//package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	j := 0
	for j < 10 {
		fmt.Println(j)
		j = j + 1
	}
	fmt.Println(j)

	z := 4
	switch z {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
	default:
		fmt.Println("Neither 0 or 1")
	}

}
