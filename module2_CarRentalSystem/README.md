- **Car Rental System in Go**
- Allows adding cars, browsing by model or price range, making, modifying, and canceling reservations
- Ensures cars cannot be double-booked and checks availability based on dates

- **Car struct**
  - Fields: `id`, `make`, `model`, `year`, `license_plate`, `price_per_day`
  - Stores all information about a car

- **allCars slice**
  - Custom type: `allCars []Car`
  - Holds all cars in the system

- **Car methods**
  - `AddCar` → pointer receiver (modifies underlying slice)
  - `SearchByModel` → value receiver (filters cars by model)
  - `SearchByPriceRange` → value receiver (filters cars by price range)

- **carsInterface**
  - Includes methods on both pointer and value receivers
  - Pointer variable assignment required to implement

- **Reservation struct**
  - Fields: `reservationid`, `carid`, `user`, `startDate`, `endDate`
  - Stores reservation details including car, user, and dates

- **allReservations slice**
  - Custom type: `allReservations []Reservation`
  - Holds all reservations

- **Reservation methods (pointer receivers)**
  - `MakeReservation` → creates new reservation if car is available
  - `ModifyReservation` → updates reservation dates if available
  - `CancelReservation` → removes reservation from slice

- **reservationsInterface**
  - Includes all pointer receiver methods
  - Pointer variable assignment required to implement

- **AvailableCars helper function**
  - Checks which cars are available for a given date range
  - Considers existing reservations to prevent overlapping bookings
  - Returns a slice of available cars
  
- **isCarAvailable function**
  - Checks if a specific car is present in slice of available cars
  - Prevents concurrent bookings by verifying availability before making or modifying a reservation