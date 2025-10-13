package main

import (
	"fmt"

	"github.com/RasagnyaVanga/carrental-db/dbpackage"
)

func main() {

	var dbstruct dbpackage.Dbstruct
	var db dbpackage.Db = &dbstruct

	db.AddCar("Hyundai", "Creta", 2020, "1345", 4000)
	db.AddCar("Honda", "City", 2022, "1445", 6000)
	db.AddCar("Kia", "Seltos", 2018, "3545", 4500)
	db.AddCar("Ford", "EcoSport", 2019, "6495", 3800)
	db.AddCar("Toyota", "Corolla", 2017, "0945", 5500)

	db.SearchByModel("Creta")

	db.SearchByPriceRange(2000, 5000)

	u := dbpackage.User{1, "John", true, "8763692534"}

	fmt.Println("Available cars from 2025-12-12 to 2025-12-12 before making reservation:")

	for _, car := range db.AvailableCars("2025-12-12", "2025-12-12") {
		fmt.Println(car)
	}

	db.MakeReservation(1, u, "2025-12-12", "2025-12-12")

	fmt.Println("Available cars: from 2025-12-12 to 2025-12-12 after making reservation:")
	for _, car := range db.AvailableCars("2025-12-12", "2025-12-12") {
		fmt.Println(car)
	}

	db.ModifyReservation(1, u, "2025-12-11", "2025-12-11")

	fmt.Println("Available cars: from 2025-12-12 to 2025-12-12 after updating reservation")
	for _, car := range db.AvailableCars("2025-12-12", "2025-12-12") {
		fmt.Println(car)
	}

	fmt.Println("Available cars: from 2025-12-11 to 2025-12-11 after updating reservation")
	for _, car := range db.AvailableCars("2025-12-11", "2025-12-11") {
		fmt.Println(car)
	}

	db.CancelReservation(1)

	fmt.Println("Available cars: from 2025-12-11 to 2025-12-11 after cancelling reservation")
	for _, car := range db.AvailableCars("2025-12-11", "2025-12-11") {
		fmt.Println(car)
	}

}
