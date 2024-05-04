package services

import "rideshare/models"

var users []models.User
var userIDCounter int

func AddUser(userDetail models.User) error {
    userDetail.ID = userIDCounter
    users = append(users, userDetail)
    userIDCounter++
    return nil
}