//package main

import "fmt"

// strut is your own data type like an object
type petRecord struct {
	name    string
	owner   string
	address string
	breed   string
	age     int
}

func (p petRecord) ageInHumanYear() int {
	return p.age * 7
}

func main() {
	fmt.Println("Hello")
	client := petRecord{"Toto", "Gale", "w233", "Dog", 9}
	fmt.Println(client.address)
	clients := make(map[int]petRecord)
	clientsId := 123
	clients[clientsId] = client
	fmt.Println(clients[clientsId].name)
	fmt.Println(clients[123].ageInHumanYear())

}
