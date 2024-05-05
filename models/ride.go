package models

type Ride struct {
	ID             int    `json:"id"`
	DriverID       int    `json:"driverId"`
	Source         string `json:"origin"`
	Destination    string `json:"destination"`
	SeatsAvailable int    `json:"seatsAvailable"`
	Passenger      int    `json:"passengers"`
	Status         string `json:"status"`
	Vehicle        string `json:"model"`
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
	UserId            int    `json:"userId"`
	Source            string `json:"source"`
	Destination       string `json:"destination"`
	Seats             int    `json:"seats"`
	SelectionStrategy string `json:"selectionStrategy"`
}

type RideManagement struct {
	RideID int `json:"rideId"`
}

type RideStats struct {
	Taken   int
	Offered int
}
