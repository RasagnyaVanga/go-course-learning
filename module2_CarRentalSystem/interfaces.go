package main

type carsInterface interface {
	AddCar(make string, model string, year int, license_plate string, price_per_day int) //on a pointer
	SearchByModel(requiredModel string) []Car                                            //on a value
	SearchByPriceRange(minPrice int, maxPrice int) []Car                                 //on a value
}

type reservationsInterface interface {
	MakeReservation(cars allCars, carid int, user user, stdate string, endate string)
	ModifyReservation(cars allCars, reservationid int, user user, nstartDate string, nendDate string)
	CancelReservation(reservationid int)
}
