package models

type Ride struct {

	ID             int      `json:"id"`
    DriverID       int      `json:"driverId"`
    Source         string   `json:"source"`
    Destination    string   `json:"destination"`
    SeatsAvailable int      `json:"seatsAvailable"`
    Passengers     []int    `json:"passengers"`
}