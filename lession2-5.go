//package main

import "fmt"

type void struct{}

var member void

type Set struct {
	members map[int]void
}

func newSet() Set {
	a := make(map[int]void)
	return Set{a}
}

func (a Set) Add(n int) {
	a.members[n] = member
}

func (a Set) Union(b Set) Set {
	c := newSet()
	for k := range a.members {
		c.members[k] = member
	}
	for k := range b.members {
		c.members[k] = member
	}
	return c
}

func (a Set) Contains(k int) bool {
	_, exists := a.members[k]
	return exists
}

func (a Set) Remove(k int) {
	delete(a.members, k)
}

func (a Set) Intersect(b Set) Set {
	c := newSet()
	for k := range a.members {
		if b.Contains(k) {
			c.members[k] = member
		}
	}
	return c
}

func (a Set) Subtract(b Set) Set {
	c := newSet()
	for k := range a.members {
		if !b.Contains(k) {
			c.members[k] = member
		}
	}
	return c
}

func (a Set) Empty() bool {
	return len(a.members) == 0
}

// Sets

func main() {

	setA := newSet()
	setB := newSet()

	setA.Add(3)
	setA.Add(4)
	setA.Add(4)
	setA.Add(5)
	setB.Add(5)
	setB.Add(56)
	fmt.Println(setA.members)
	fmt.Println("End")
	fmt.Println(setA.Union(setB).members)
	setC := newSet()
	fmt.Println("C Empty", setC.Empty())
	fmt.Println("A Empty", setA.Empty())
	fmt.Println("Interset", setA.Intersect(setB).members)
	setD := newSet()
	setD.Add(4)
	setD.Add(5)
	fmt.Println(setD.members)
	setD.Remove(4)
	fmt.Println(setD.members)
	fmt.Println(setD.Subtract(setB).members)

}
