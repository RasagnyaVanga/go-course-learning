package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func recoverPanic() {
	if r := recover(); r != nil {
		fmt.Println("Unexpected error occurred:", r)
	}
}

func main() {
	defer recoverPanic()

	stringOfNumbers := flag.String("numbers", "", "numbers seperated by comma in a string")
	flag.Parse()

	if *stringOfNumbers == "" {
		fmt.Println("Error: No input provided. Please use -numbers flag with values (e.g. -numbers=1,2,3)")
		os.Exit(1)
	}

	numbersStringArray := strings.Split(*stringOfNumbers, ",")
	var numbersIntegerArray []int
	for _, num := range numbersStringArray {
		num = strings.TrimSpace(num)
		n, err := strconv.Atoi(num)
		if err != nil {
			fmt.Printf("Skipping invalid input '%s': %v\n", num, err)
			continue
		}
		numbersIntegerArray = append(numbersIntegerArray, n)
	}
	if len(numbersIntegerArray) == 0 { //string given through command line, but no valid numbers out of them
		panic("No valid numbers were provided. Cannot continue.")
	}

	var s int
	for _, n := range numbersIntegerArray {
		s += n
	}
	fmt.Println("sum of numbers is ", s)
}
