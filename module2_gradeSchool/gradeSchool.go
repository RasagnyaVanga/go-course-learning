package main

import (
	"fmt"
	"slices"
	"sort"
)

type Record struct {
	data map[int][]string
}

func (r *Record) Add(grade int, newStudent string, allStudents []string) {
	if slices.Contains(allStudents, newStudent) {
		return
	}
	r.data[grade] = append(r.data[grade], newStudent)
	sort.Strings(r.data[grade])
}
func (r *Record) getStudentsInGrade(grade int) []string {
	return r.data[grade]
}
func (r *Record) All() []string {
	var result []string
	var grades []int

	for g := range r.data {
		grades = append(grades, g)
	}
	sort.Ints(grades)

	for _, g := range grades {
		result = append(result, r.data[g]...)
	}
	return result
}
func main() {
	r := Record{data: make(map[int][]string)}
	r.Add(1, "Alex", r.All())
	r.Add(2, "Alex", r.All())
	r.Add(1, "Peter", r.All())
	r.Add(2, "Zoe", r.All())
	r.Add(5, "Jim", r.All())
	r.Add(2, "Anna", r.All())
	r.Add(5, "Zoe", r.All())
	fmt.Println("Students in grade 1", r.getStudentsInGrade(1))
	fmt.Println("Students in grade 2", r.getStudentsInGrade(2))
	fmt.Println("Students in grade 5", r.getStudentsInGrade(5))

	fmt.Println("All Students :", r.All())
}
