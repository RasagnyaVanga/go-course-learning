package main

import (
	"fmt"

	"github.com/RasagnyaVanga/carrental/cars"
	"github.com/RasagnyaVanga/carrental/interfaces"
	res "github.com/RasagnyaVanga/carrental/reservations"
	"github.com/RasagnyaVanga/carrental/users"
)

func main() {
	var reservations res.Reservations
	var reservationHandler interfaces.ReservationHandler = &reservations
	var cars cars.Cars
	var carService interfaces.CarService = &cars

	carService.AddCar("Hyundai", "Creta", 2020, "1345", 4000)
	carService.AddCar("Honda", "City", 2022, "1445", 6000)
	carService.AddCar("Kia", "Seltos", 2018, "3545", 4500)
	carService.AddCar("Ford", "EcoSport", 2019, "6495", 3800)
	carService.AddCar("Toyota", "Corolla", 2017, "0945", 5500)

	carService.SearchByModel("Creta")

	carService.SearchByPriceRange(2000, 5000)

	u := users.User{1, "John", true, "8763692534"}

	fmt.Println("Available cars from 2025-12-12 to 2025-12-12 before making reservation:")
	for _, car := range res.AvailableCars(cars, reservations, "2025-12-12", "2025-12-12") {
		fmt.Println(car)
	}

	reservationHandler.MakeReservation(cars, 1, u, "2025-12-12", "2025-12-12")

	fmt.Println("Available cars: from 2025-12-12 to 2025-12-12 after making reservation:")
	for _, car := range res.AvailableCars(cars, reservations, "2025-12-12", "2025-12-12") {
		fmt.Println(car)
	}

	reservationHandler.ModifyReservation(cars, 1, u, "2025-12-11", "2025-12-11")

	fmt.Println("Available cars: from 2025-12-12 to 2025-12-12 after updating reservation")
	for _, car := range res.AvailableCars(cars, reservations, "2025-12-12", "2025-12-12") {
		fmt.Println(car)
	}

	fmt.Println("Available cars: from 2025-12-11 to 2025-12-11 after updating reservation")
	for _, car := range res.AvailableCars(cars, reservations, "2025-12-11", "2025-12-11") {
		fmt.Println(car)
	}

	reservationHandler.CancelReservation(1)

	fmt.Println("Available cars: from 2025-12-11 to 2025-12-11 after cancelling reservation")
	for _, car := range res.AvailableCars(cars, reservations, "2025-12-11", "2025-12-11") {
		fmt.Println(car)
	}

}
