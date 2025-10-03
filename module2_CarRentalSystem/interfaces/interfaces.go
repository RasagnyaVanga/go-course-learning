package interfaces

import (
	"github.com/RasagnyaVanga/carrental/cars"
	"github.com/RasagnyaVanga/carrental/users"
)

type CarService interface {
	AddCar(make string, model string, year int, license_plate string, price_per_day int) //on a pointer
	SearchByModel(requiredModel string) []cars.Car                                       //on a value
	SearchByPriceRange(minPrice int, maxPrice int) []cars.Car                            //on a value
}

type ReservationHandler interface {
	MakeReservation(cars cars.Cars, carid int, user users.User, stdate string, endate string)
	ModifyReservation(cars cars.Cars, reservationid int, user users.User, nstartDate string, nendDate string)
	CancelReservation(reservationid int)
}
