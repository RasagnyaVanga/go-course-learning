package main

import "testing"

// func TestFooer(t *testing.T) {
// 	got := Fooer(3)
// 	exp := "isDivisible"

// 	if got != exp {
// 		t.Errorf("Result was incorrect, got: %s, want: %s.", got, exp)
// 	}
// }

func TestFooerTableDriven(t *testing.T) {
	var tests = []struct {
		name  string
		input int
		exp   string
	}{
		{"Divisible", 9, "isDivisible"},
		{"Not Divisible", 10, "10"},
		{"0 Divisible", 0, "isDivisible"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := Fooer(tt.input)
			if ans != tt.exp {
				t.Errorf("got %s, want %s", ans, tt.exp)
			}
		})
	}
}

func TestFooer(t *testing.T) {
	t.Run("3 isDivisible", func(t *testing.T) {
		t.Parallel()
		got := Fooer(3)
		exp := "isDivisible"
		if got != exp {
			t.Errorf("Result was incorrect, got: %s, want: %s.", got, exp)
		}
	})
	t.Run("7 isNotDivisible", func(t *testing.T) {
		t.Parallel()
		got := Fooer(7)
		exp := "7"
		if got != exp {
			t.Errorf("Result was incorrect, got: %s, want: %s.", got, exp)
		}
	})
}

// func TestAdd(t *testing.T) {
// 	got := Add(2, 3)
// 	exp := 5
// 	if got != exp {
// 		t.Errorf("got %d expected %d", got, exp)
// 	}
// }

func TestAdd(t *testing.T) {
	t.Run("both positive", func(t *testing.T) {
		got := Add(2, 3)
		exp := 5
		if got != exp {
			t.Errorf("got %d expected %d", got, exp)
		}
	})
	t.Run("both negative", func(t *testing.T) {
		got := Add(-2, -3)
		exp := -5
		if got != exp {
			t.Errorf("got %d expected %d", got, exp)
		}
	})
}

func TestAddTableDriven(t *testing.T) {
	var tests = []struct {
		name string
		a    int
		b    int
		exp  int
	}{
		{"Both positive", 2, 3, 5},
		{"greater positive- lesser negative", -2, 3, 1},
		{"greater negative- lesser positive", 2, -3, -1},
		{"both negative", -2, -3, -5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Add(tt.a, tt.b)
			if got != tt.exp {
				t.Errorf("got %d expected %d", got, tt.exp)
			}
		})
	}
}

func TestIsEven(t *testing.T) {
	var tests = []struct {
		name  string
		input int
		exp   bool
	}{
		{"even", 2, true},
		{"odd", 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsEven(tt.input)
			if got != tt.exp {
				t.Errorf("got %v exp %v", got, tt.exp)
			}
		})
	}
}

func TestDeposit(t *testing.T) {
	var tests = []struct {
		name   string
		amount float64
		exp    float64
	}{
		{"positive", 50, 50},
		{"negative", -50, -50},
		{"0 value", 0, -0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Account{Balance: 0}
			d.Deposit(tt.amount)
			if d.Balance != tt.exp {
				t.Errorf("got %f expected %f", d.Balance, tt.exp)
			}
		})
	}
}
