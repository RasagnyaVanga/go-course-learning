package main

import "fmt"

type Reservation struct {
	reservationid int
	carid         int
	user          user
	startDate     string
	endDate       string
}

type allReservations []Reservation //custom type allReservations

var nextReservationID = 1

func (r *allReservations) MakeReservation(cars allCars, carid int, user user, stdate string, endate string) {
	if !isCarAvailable(AvailableCars(cars, *r, stdate, endate), carid) {
		fmt.Printf("This car is not available to reserve on %s to %s\n", stdate, endate)
		return
	}
	res := Reservation{nextReservationID, carid, user, stdate, endate}
	res.reservationid = nextReservationID
	nextReservationID++
	*r = append(*r, res)
	fmt.Printf("Reservation made successfully reservation-id : %d, car-id:%d, user-id:%d from %s to %s\n", res.reservationid, res.carid, res.user.userid, res.startDate, res.endDate)
}

func (r *allReservations) ModifyReservation(cars allCars, id int, user user, newStartDate string, newEndDate string) {
	fmt.Println("MODIFYING RESERVATION")
	fmt.Println("Old reservation", r)
	for i, rec := range *r {
		if rec.reservationid == id {
			if !isCarAvailable(AvailableCars(cars, *r, newStartDate, newEndDate), rec.carid) {
				fmt.Println("Car is not available on newdates")
				return
			}
			(*r)[i].startDate = newStartDate
			(*r)[i].endDate = newEndDate
		}
	}

	fmt.Println("Updated reservation", r)
}

func (r *allReservations) CancelReservation(id int) {
	for i, rec := range *r {
		if rec.reservationid == id {
			fmt.Printf("CANCELLING RESERVATIONreservation-id : %d, car-id:%d, user-id:%d from %s to %s\n", rec.reservationid, rec.carid, rec.user.userid, rec.startDate, rec.endDate)
			*r = append((*r)[:i], (*r)[i+1:]...)
		}
	}
}
