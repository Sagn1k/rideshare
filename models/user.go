package models

type User struct {
	ID          int    `json:"id"`
    Name        string `json:"name"`
    Email       string `json:"email"`
    PhoneNumber string `json:"phoneNumber"`
	//TODO: Make Role as enum
    Role        string `json:"role"` //Passenger or Driver
}
