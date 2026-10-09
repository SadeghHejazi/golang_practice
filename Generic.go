package main

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[G any] struct {
	next *List[G]
	val  G
}

func (L *List[G]) Push(val G) *List[G] {
	newnode := &List[G]{
		val:  val,
		next: L,
	}
	return newnode
}
func (L *List[G]) Print() {
	for node := L; node != nil; node = node.next {
		fmt.Println(node.val)
	}
}

// Index returns the index of x in s, or -1 if not found.Generic one
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "baz"))

	//main for generic two
	var list *List[int]

	list = list.Push(10)
	list = list.Push(20)
	list = list.Push(30)

	list.Print()

}
