package main

type Car struct {
	id            int
	make          string
	model         string
	year          int
	license_plate string
	price_per_day int
}

type allCars []Car //custom type - allCars

var nextCarId = 1

func (cars *allCars) AddCar(make string, model string, year int, license_plate string, price_per_day int) {
	newCar := Car{nextCarId, make, model, year, license_plate, price_per_day}
	nextCarId++
	//todo - are pointers required? //yes inorder to change the underlying values of slice, we need pointers
	*cars = append(*cars, newCar)
}

func (cars allCars) SearchByModel(requiredModel string) []Car {
	var resultcars []Car
	for _, c := range cars {
		if c.model == requiredModel {
			resultcars = append(resultcars, c)
		}
	}
	return resultcars
}

func (cars allCars) SearchByPriceRange(minPrice int, maxPrice int) []Car {
	var resultcars []Car
	for _, c := range cars {
		if c.price_per_day >= minPrice && c.price_per_day <= maxPrice {
			resultcars = append(resultcars, c)
		}
	}
	return resultcars
}
