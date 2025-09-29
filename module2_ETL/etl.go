package main

import (
	"fmt"
	"strings"
)

func main() {
	onetomany := make(map[int][]string)
	onetomany[1] = []string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}
	onetomany[2] = []string{"D", "G"}
	onetomany[3] = []string{"B", "C", "M", "P"}
	onetomany[4] = []string{"F", "H", "V", "W", "Y"}
	onetomany[5] = []string{"K"}
	onetomany[8] = []string{"J", "X"}
	onetomany[10] = []string{"Q", "Z"}
	// var onetoone map[string]int
	onetoone := make(map[string]int)
	for key, values := range onetomany {
		for _, val := range values {
			onetoone[strings.ToLower(val)] = key
		}
	}
	fmt.Println(onetoone)
}
