package reservations

import (
	"fmt"
	"os"
	"time"

	"github.com/RasagnyaVanga/carrental/users"

	carspackage "github.com/RasagnyaVanga/carrental/cars"
)

type Reservation struct {
	id        int
	carid     int
	user      users.User
	startDate string
	endDate   string
}

type Reservations []Reservation //custom type Reservations

var nextReservationID = 1

func (r *Reservations) MakeReservation(cars carspackage.Cars, carid int, user users.User, startDate string, endDate string) {
	if !carspackage.IsCarAvailable(AvailableCars(cars, *r, startDate, endDate), carid) {
		fmt.Printf("This car is not available to reserve on %s to %s\n", startDate, endDate)
		return
	}
	res := Reservation{nextReservationID, carid, user, startDate, endDate}
	res.id = nextReservationID
	nextReservationID++
	*r = append(*r, res)
	fmt.Printf("Reservation made successfully reservation-id : %d, car-id:%d, user-id:%d from %s to %s\n", res.id, res.carid, res.user.Id, res.startDate, res.endDate)
}

func (r *Reservations) ModifyReservation(cars carspackage.Cars, id int, user users.User, startDate string, endDate string) {
	fmt.Println("MODIFYING RESERVATION")
	fmt.Println("Old reservation", r)
	for i, rec := range *r {
		if rec.id == id {
			if !carspackage.IsCarAvailable(AvailableCars(cars, *r, startDate, endDate), rec.carid) {
				fmt.Println("Car is not available on newdates")
				return
			}
			(*r)[i].startDate = startDate
			(*r)[i].endDate = endDate
		}
	}

	fmt.Println("Updated reservation", r)
}

func (r *Reservations) CancelReservation(id int) {
	for i, rec := range *r {
		if rec.id == id {
			fmt.Printf("CANCELLING RESERVATIONreservation-id : %d, car-id:%d, user-id:%d from %s to %s\n", rec.id, rec.carid, rec.user.Id, rec.startDate, rec.endDate)
			*r = append((*r)[:i], (*r)[i+1:]...)
		}
	}
}

func AvailableCars(cars carspackage.Cars, reservations Reservations, reqStartDate string, reqEndDate string) []carspackage.Car {
	format := "2006-01-02"
	queryStart, err := time.Parse(format, reqStartDate)
	if err != nil {
		fmt.Println("Invalid date format:", reqStartDate)
		os.Exit(1)
	}

	queryEnd, err := time.Parse(format, reqEndDate)
	if err != nil {
		fmt.Println("Invalid date format:", reqEndDate)
		os.Exit(1)
	}

	reserved := make(map[int]bool)
	for _, res := range reservations {
		start, err := time.Parse(format, res.startDate)
		if err != nil {
			fmt.Println("Invalid reservation date:", res.startDate)
			os.Exit(1)
		}
		end, err := time.Parse(format, res.endDate)
		if err != nil {
			fmt.Println("Invalid reservation date:", res.endDate)
			os.Exit(1)
		}

		if !(queryEnd.Before(start) || queryStart.After(end)) {
			reserved[res.carid] = true
		}
	}

	var availablecars []carspackage.Car

	for _, car := range cars {
		if reserved[car.Id] {
			continue
		}
		availablecars = append(availablecars, car)
	}
	return availablecars
}
