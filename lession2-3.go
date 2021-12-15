//package main

import "fmt"

func main() {
	m := make(map[string]string)
	m["First Name"] = "Will"
	m["Last Name"] = "Doe"
	m["Address"] = "123 Test St"
	m["City"] = "Tampa"
	m["Country"] = "Pinellas"
	m["Zip"] = "33756"

	fmt.Println(m["First Name"])

	fmt.Println(m["Address"] + "," + m["City"])
	delete(m, "First Name")
	fmt.Println(m["First Name"])
	fmt.Println(m["Middle Name"])

}
