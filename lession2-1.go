//package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println("a:", a)

	a[2] = 6
	fmt.Println("a:", a)

	groceryList := [4]string{"apples", "carrots", "flour", "chocolate"}

	for i, s := range groceryList {
		fmt.Println(i, s)
	}

	fmt.Println(len(groceryList))

	hardwareStore := [4]string{"hammer", "nails", "tape", "paint"}
	clothingStore := [4]string{"shirt", "socks", "belt", "cuff links"}

	shoppingLists := [3][4]string{groceryList, hardwareStore, clothingStore}

	for i, list := range shoppingLists {
		fmt.Println(i, list)
		for j, item := range list {
			fmt.Println(j, item)
		}
	}
}
