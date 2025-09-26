package main

import (
	"fmt"
	"time"
)

func isCarAvailable(availableCars []Car, carid int) bool {
	for _, c := range availableCars {
		if c.id == carid {
			return true
		}
	}
	return false
}

func AvailableCars(cars allCars, reservations allReservations, stdate string, endate string) []Car {
	format := "2006-01-02"
	queryStart, err1 := time.Parse(format, stdate)
	queryEnd, err2 := time.Parse(format, endate)
	if err1 != nil || err2 != nil {
		fmt.Println("Invalid date format:", stdate)
	}

	reserved := make(map[int]bool)
	for _, res := range reservations {
		start, err1 := time.Parse(format, res.startDate)
		end, err2 := time.Parse(format, res.endDate)
		if err1 != nil || err2 != nil {
			fmt.Println("Invalid reservation date:", res.startDate, "-", res.endDate)
			continue
		}

		if !(queryEnd.Before(start) || queryStart.After(end)) {
			reserved[res.carid] = true
		}
	}

	var availablecars []Car

	for _, car := range cars {
		if reserved[car.id] {
			continue
		}
		availablecars = append(availablecars, car)
	}
	return availablecars
}

func main() {
	var reservations allReservations
	var reservationsInterface reservationsInterface = &reservations
	var cars allCars
	var carsInterface carsInterface = &cars

	carsInterface.AddCar("Hyundai", "Creta", 2020, "1345", 4000)
	carsInterface.AddCar("Honda", "City", 2022, "1445", 6000)
	carsInterface.AddCar("Kia", "Seltos", 2018, "3545", 4500)
	carsInterface.AddCar("Ford", "EcoSport", 2019, "6495", 3800)
	carsInterface.AddCar("Toyota", "Corolla", 2017, "0945", 5500)

	carsInterface.SearchByModel("Creta")

	carsInterface.SearchByPriceRange(2000, 5000)

	u := user{1, "John", true, "8763692534"}

	fmt.Println("Available cars from 2025-12-12 to 2025-12-12 before making reservation:")
	for _, car := range AvailableCars(cars, reservations, "2025-12-12", "2025-12-12") {
		fmt.Println(car)
	}

	reservationsInterface.MakeReservation(cars, 1, u, "2025-12-12", "2025-12-12")

	fmt.Println("Available cars: from 2025-12-12 to 2025-12-12 after making reservation:")
	for _, car := range AvailableCars(cars, reservations, "2025-12-12", "2025-12-12") {
		fmt.Println(car)
	}

	reservationsInterface.ModifyReservation(cars, 1, u, "2025-12-11", "2025-12-11")

	fmt.Println("Available cars: from 2025-12-12 to 2025-12-12 after updating reservation")
	for _, car := range AvailableCars(cars, reservations, "2025-12-12", "2025-12-12") {
		fmt.Println(car)
	}

	fmt.Println("Available cars: from 2025-12-11 to 2025-12-11 after updating reservation")
	for _, car := range AvailableCars(cars, reservations, "2025-12-11", "2025-12-11") {
		fmt.Println(car)
	}

	reservationsInterface.CancelReservation(1)

	fmt.Println("Available cars: from 2025-12-11 to 2025-12-11 after cancelling reservation")
	for _, car := range AvailableCars(cars, reservations, "2025-12-11", "2025-12-11") {
		fmt.Println(car)
	}

	// fmt.Println(`Enter the start date you want to book car in format "yyyy-mm-dd"`)
	// var startDate string
	// fmt.Scanln(&startDate)
	// fmt.Println(`Enter the end date you want to book car in format "yyyy-mm-dd"`)
	// var endDate string
	// fmt.Scanln(&endDate)
	// fmt.Printf("Available cars from %s to %s :", startDate, endDate)
	// for _, car := range AvailableCars(startDate, endDate) {
	// 	fmt.Printf("car id: %d, Made by: %s, Model: %s, year:%d, license_plate: %s, price_per_day:%d",
	// 		car.id, car.make, car.model, car.year, car.license_plate, car.price_per_day)
	// }
	// fmt.Println("Enter the carid you want to reserve")
	// var carid int
	// fmt.Scanln(&carid)
	// var presentBookingCar car
	// for _, car := range cars {
	// 	if car.id == carid {
	// 		presentBookingCar = car
	// 		break
	// 	}
	// }
	// fmt.Println("Confirm booking car ", presentBookingCar)
	// fmt.Printf("from %s to %s yes or no?", startDate, endDate)
	// var bookConfirmation string
	// fmt.Scanln(&bookConfirmation)
	// if bookConfirmation == "yes" {
	// 	reservations.makeReservation(carid, u, startDate, endDate)
	// } else {
	// 	return
	// }
	// fmt.Println("Your reservations: ")
	// for _, res := range reservations {
	// 	if res.user.userid == 1 {
	// 		fmt.Println(res)
	// 	}
	// }
	// fmt.Println("Do you want to modify your booking ( change dates? )")
	// var modifyConfirmation string
	// fmt.Scanln(&modifyConfirmation)
	// if modifyConfirmation == "yes" {
	// 	fmt.Println("Enter reservation id")
	// 	var resid int
	// 	fmt.Scanln(&resid)
	// 	fmt.Println(`Enter the start date you want to book car in format "yyyy-mm-dd"`)
	// 	var nstartDate string
	// 	fmt.Scanln(&nstartDate)
	// 	fmt.Println(`Enter the end date you want to book car in format "yyyy-mm-dd"`)
	// 	var nendDate string
	// 	fmt.Scanln(&nendDate)

	// 	reservations.modifyReservation(resid, u, nstartDate, nendDate)
	// } else {
	// 	return
	// }
	// fmt.Println("Do you want to cancel any booking ( change dates? )")
	// var cancelConfirmation string
	// fmt.Scanln(&cancelConfirmation)
	// if cancelConfirmation == "yes" {
	// 	fmt.Println("Enter reservation id")
	// 	var resid int
	// 	fmt.Scanln(&resid)
	// 	reservations.cancelReservation(resid)
	// } else {
	// 	return
	// }

}
