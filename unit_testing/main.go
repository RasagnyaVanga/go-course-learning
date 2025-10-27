package main

import "strconv"

// basic from an article
func Fooer(input int) string {
	isDivisible := input%3 == 0
	if isDivisible {
		return "isDivisible"
	}
	return strconv.Itoa(input)
}

// from mentor: 1.Basic Function Testing
func Add(a, b int) int {
	return a + b
}

// from mentor: 2. Table-Driven Tests
func IsEven(n int) bool {
	return n%2 == 0
}

// from mentor:3. Testing Methods on Structs
type Account struct {
	Balance float64
}

func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}
