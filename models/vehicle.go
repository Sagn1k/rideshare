package models

type Vehicle struct {
	ID           int    `json:"id"`
    UserID       int    `json:"userId"`
    Brand        string `json:"brand"`
    Model        string `json:"model"`
    Year         int    `json:"year"`
    LicensePlate string `json:"licensePlate"`
}
