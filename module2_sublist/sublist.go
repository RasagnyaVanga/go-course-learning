package main

import (
	"fmt"
)

func isEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true

}
func isSublist(a, b []int) bool {
	if len(a) > len(b) {
		return false
	}
	ai := 0
	bi := 0

	for bi < len(b) && ai < len(a) {
		if a[ai] == b[bi] {
			ai++
			bi++
		} else {
			if ai > 0 {
				ai = 0
			} else {
				bi++
			}
		}
	}
	return ai == len(a)
}
func main() {
	A := []int{1, 2, 3, 2, 3, 4, 4}
	B := []int{2, 3, 4}
	if isEqual(A, B) {
		fmt.Println("List A is equal to list B")
	} else if isSublist(B, A) {
		fmt.Println("List A contains list B (A is a superlist of B)")
	} else if isSublist(A, B) {
		fmt.Println("List A is contained by list B (A is a sublist of B)")
	} else {
		fmt.Println("Lists A and B are unequal")
	}
}
