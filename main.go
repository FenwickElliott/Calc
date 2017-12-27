package main

import (
	"fmt"
	"os"
)

type clause struct {
	num      int
	operator string
}

func main() {
	input := os.Args[1:][0]
	clauses := []clause{}

	for _, v := range input {
		if 48 <= v && v <= 57 {
			clauses = append(clauses, clause{num: int(v - 48)})
			// } else if opperator {
			// handle opperators
		} else {
			fmt.Println("Error: That was not a mathmatical character")
			return
		}
	}

	fmt.Println(clauses)
}

func multiply(nums ...int) int {
	res := 1
	for _, num := range nums {
		res *= num
	}
	return res
}

// */+-
