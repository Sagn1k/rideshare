package models

type Ride struct {
	ID             int    `json:"id"`
	DriverID       int    `json:"driverId"`
	Source         string `json:"origin"`
	Destination    string `json:"destination"`
	SeatsAvailable int    `json:"seatsAvailable"`
	Passengers     []int  `json:"passengers"`
	Status         string `json:"status"`
	Active         bool   `json:"isActive"`
}

type RideOfferRequest struct {
	DriverID       int    `json:"driverId"`
	Source         string `json:"origin"`
	Destination    string `json:"destination"`
	SeatsAvailable int    `json:"seatsAvailable"`
	Vehicle        string `json:"vehicle"`
}

type RideSelection struct {
	Source            string `json:"source"`
	Destination       string `json:"destination"`
	Seats             int    `json:"seats"`
	SelectionStrategy string `json:"selectionStrategy"`
}

type RideManagement struct {
	RideID int `json:"rideId"`
}
