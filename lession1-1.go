// package main  //comment out if have more than one main in project

import "fmt"

func main() {
	//fmt.Println("Helloo" + " world")
	//fmt.Println("Hello", 1+2)
	//fmt.Println("bool ", !true)
	// DeMorgan's Law
	fmt.Println(!(true || false) == (true && false))
	fmt.Println(!(true && false) == (true || false))
}
