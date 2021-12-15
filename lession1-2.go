//package main

import "fmt"

const PI = 3.14159

func main() {
	name := "William"
	const greeting = "HELLO,"
	fmt.Println(greeting, name, PI)
	radius := 8.0
	fmt.Println("circumference:", PI*(radius+radius))
	fmt.Println("area:", PI*(radius*radius))
	radius = 5
	fmt.Println("circumference:", PI*(radius+radius))
	fmt.Println("area:", PI*(radius*radius))
}
