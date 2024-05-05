package models

type Ride struct {

	ID           int
    DriverID     int
    Source       string
    Destination  string
    Available    int // Number of available seats
    Active       bool
}