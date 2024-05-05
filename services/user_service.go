package services

import (
	"errors"
	"rideshare/models"

	"go.uber.org/zap"
)

var users []models.User
var userIDCounter int

type UserService struct {
    logger *zap.Logger
}

// NewUserService creates a new instance of UserService with the provided logger
func NewUserService(logger *zap.Logger) *UserService {
    return &UserService{logger: logger}
}

func (s *UserService) AddUser(userDetail models.User) error {
	userDetail.ID = userIDCounter
	users = append(users, userDetail)
	userIDCounter++

	s.logger.Info("Adding user", zap.String("name", userDetail.Name))

	return nil
}

// GetUser retrieves a user by ID from the in-memory storage
func (s *UserService) GetUser(userID int) (*models.User, error) {

	s.logger.Info("Getting user", zap.Int("userID", userID))

	// Loop through the users slice to find the user with the given ID
	for _, user := range users {
		if user.ID == userID {
			// If the user is found, return the user
			return &user, nil
		}
	}
	// If the user is not found, return an error
	return nil, errors.New("user not found")
}

// GetAllUsers retrieves all users from the in-memory storage
func (s *UserService) GetAllUsers() ([]models.User, error) {

	s.logger.Info("Getting all users") 
	return users, nil
}
