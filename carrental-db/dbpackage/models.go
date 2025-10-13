package dbpackage

type Car struct {
	Id           int
	make         string
	model        string
	year         int
	licensePlate string
	pricePerDay  int
}
type Reservation struct {
	reservationid int
	carid         int
	user          User
	startDate     string
	endDate       string
}

type User struct {
	Id          int
	Username    string
	License     bool
	Phonenumber string
}
type cars []Car
type reservations []Reservation

type Dbstruct struct {
	cars         []Car
	reservations []Reservation
}

type Db interface {
	AddCar(make string, model string, year int, license_plate string, price_per_day int) //on a pointer
	SearchByModel(requiredModel string) []Car                                            //on a value
	SearchByPriceRange(minPrice int, maxPrice int) []Car                                 //on a value
	MakeReservation(carid int, user User, stdate string, endate string)
	ModifyReservation(reservationid int, user User, nstartDate string, nendDate string)
	CancelReservation(reservationid int)
	AvailableCars(stdate string, endate string) []Car
}
