- **Car Rental System in Go**
  - Allows adding cars, browsing by model or price range, making, modifying, and canceling reservations
  - Ensures cars cannot be double-booked and checks availability based on dates

- **Cars package**
  - **Car struct**
    - Fields: `Id`, `make`, `model`, `year`, `licensePlate`, `pricePerDay`
    - Stores all information about a car

  - **Cars slice**
    - Custom type: `Cars []Car`
    - Holds all cars in the system

  - **Car methods**
    - `AddCar` → pointer receiver (modifies underlying slice)
    - `SearchByModel` → value receiver (filters cars by model)
    - `SearchByPriceRange` → value receiver (filters cars by price range)

  - **IsCarAvailable function**
    - Checks if a specific car is present in slice of available cars
    - Prevents concurrent bookings by verifying availability before making or modifying a reservation

- **Reservations package**
  - **Reservation struct**
    - Fields: `id`, `carid`, `user`, `startDate`, `endDate`
    - Stores reservation details including car, user, and dates

  - **Reservations slice**
    - Custom type: `Reservations []Reservation`
    - Holds all reservations

  - **Reservation methods (pointer receivers)**
    - `MakeReservation` → creates new reservation if car is available
    - `ModifyReservation` → updates reservation dates if available
    - `CancelReservation` → removes reservation from slice

  - **AvailableCars function**
    - Checks which cars are available for a given date range
    - Considers existing reservations to prevent overlapping bookings
    - Returns a slice of available cars

- **Users package**
  - **User struct**
    - Fields: `Id`, `Username`, `License`, `Phonenumber`
    - Stores user details and license information

- **Interfaces package**
  - **CarService interface**
    - Includes methods on both pointer and value receivers
    - Pointer variable assignment required to implement

  - **ReservationHandler interface**
    - Includes all pointer receiver methods
    - Pointer variable assignment required to implement

- **Main functionality**
  - Demonstrates adding cars, searching by model and price range
  - Shows making, modifying, and canceling reservations
  - Displays available cars before and after reservation operations
