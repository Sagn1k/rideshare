package models

type Vehicle struct {
	ID           int    `json:"id"`
    UserID       int    `json:"userId"`
    Model        string `json:"model"`
    LicensePlate string `json:"licensePlate"`
}
