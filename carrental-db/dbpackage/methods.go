package dbpackage

import (
	"fmt"
	"time"
)

//cars

var nextCarId = 1

func (db *Dbstruct) AddCar(make string, model string, year int, licensePlate string, pricePerDay int) {
	newCar := Car{nextCarId, make, model, year, licensePlate, pricePerDay}
	nextCarId++
	(*db).cars = append((*db).cars, newCar)
}

func (db *Dbstruct) SearchByModel(requiredModel string) []Car {
	var results []Car
	for _, c := range db.cars {
		if c.model == requiredModel {
			results = append(results, c)
		}
	}
	return results
}

func (db *Dbstruct) SearchByPriceRange(minPrice int, maxPrice int) []Car {
	var results []Car
	for _, c := range db.cars {
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

//reservations

var nextReservationID = 1

func (db *Dbstruct) MakeReservation(carid int, user User, stdate string, endate string) {
	if !IsCarAvailable(db.AvailableCars(stdate, endate), carid) {
		fmt.Printf("This car is not available to reserve on %s to %s\n", stdate, endate)
		return
	}
	res := Reservation{nextReservationID, carid, user, stdate, endate}
	res.reservationid = nextReservationID
	nextReservationID++
	(*db).reservations = append((*db).reservations, res)
	fmt.Printf("Reservation made successfully reservation-id : %d, car-id:%d, user-id:%d from %s to %s\n", res.reservationid, res.carid, res.user.Id, res.startDate, res.endDate)
}

func (db *Dbstruct) ModifyReservation(id int, user User, newStartDate string, newEndDate string) {
	fmt.Println("MODIFYING RESERVATION")
	fmt.Println("Old reservation", db.reservations)
	for i, rec := range (*db).reservations {
		if rec.reservationid == id {
			if !IsCarAvailable(db.AvailableCars(newStartDate, newEndDate), rec.carid) {
				fmt.Println("Car is not available on newdates")
				return
			}
			(*db).reservations[i].startDate = newStartDate
			(*db).reservations[i].endDate = newEndDate
		}
	}

	fmt.Println("Updated reservation", db.reservations)
}

func (db *Dbstruct) CancelReservation(id int) {
	for i, rec := range (*db).reservations {
		if rec.reservationid == id {
			fmt.Printf("CANCELLING RESERVATIONreservation-id : %d, car-id:%d, user-id:%d from %s to %s\n", rec.reservationid, rec.carid, rec.user.Id, rec.startDate, rec.endDate)
			(*db).reservations = append((*db).reservations[:i], (*db).reservations[i+1:]...)
		}
	}
}

func (db *Dbstruct) AvailableCars(stdate string, endate string) []Car {
	format := "2006-01-02"
	queryStart, err1 := time.Parse(format, stdate)
	queryEnd, err2 := time.Parse(format, endate)
	if err1 != nil || err2 != nil {
		fmt.Println("Invalid date format:", stdate)
	}

	reserved := make(map[int]bool)
	for _, res := range db.reservations {
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

	for _, car := range db.cars {
		if reserved[car.Id] {
			continue
		}
		availablecars = append(availablecars, car)
	}
	return availablecars
}
