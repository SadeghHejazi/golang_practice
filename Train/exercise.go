package main

import (
	"fmt"
)

func main() {
	var input string
	fmt.Scan(&input)
	var newstring string
	for _, v := range input {
		if v != ' ' {
			newstring += string(v)
		}
		input = newstring
	}

	fmt.Println(input)
}
