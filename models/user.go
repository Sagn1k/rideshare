package models

type User struct {
	ID       int
    Name     string
	//TODO: change role to enum
    Role     string // Driver or Passenger
    Vehicles []Vehicle
}
