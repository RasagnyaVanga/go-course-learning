package cars

import "fmt"

type Car struct {
	Id           int
	make         string
	model        string
	year         int
	licensePlate string
	pricePerDay  int
}

type Cars []Car //custom type - Cars

var nextCarId = 1

func (cars *Cars) AddCar(make string, model string, year int, licensePlate string, pricePerDay int) {
	newCar := Car{nextCarId, make, model, year, licensePlate, pricePerDay}
	nextCarId++
	*cars = append(*cars, newCar)
}

func (cars Cars) SearchByModel(requiredModel string) []Car {
	var results []Car
	if requiredModel == "" {
		fmt.Println("Enter valid model name of car")
		return results
	}
	for _, c := range cars {
		if c.model == requiredModel {
			results = append(results, c)
		}
	}
	return results
}

func (cars Cars) SearchByPriceRange(minPrice int, maxPrice int) []Car {
	var results []Car

	if minPrice < 0 || maxPrice < 0 {
		fmt.Println("price should be positive")
		return results
	}
	if minPrice > maxPrice {
		fmt.Println("Enter valid minimum and maximum prices")
		return results
	}
	for _, c := range cars {
		if c.pricePerDay >= minPrice && c.pricePerDay <= maxPrice {
			results = append(results, c)
		}
	}
	return results
}

func IsCarAvailable(availableCars []Car, carid int) bool {
	for _, c := range availableCars {
		if c.Id == carid {
			return true
		}
	}
	return false
}
