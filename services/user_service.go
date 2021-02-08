package services

import (
	"errors"
	"github.com/benbarron/golang-auth-server/database"
	"gorm.io/gorm"
)

type UserService struct {
	repo *gorm.DB
	logger *LoggingService
}

func NewUserService() *UserService {
	return &UserService{
		logger: NewLoggingService("UserService"),
		repo:   database.GetDatabase(),
	}
}

func (s *UserService) FindAll() ([]database.User, error) {
	var users []database.User
	s.repo.Find(&users)
	return users, nil
}

func (s *UserService) CreateUser(user *database.User) error {
	_ = user.HashPassword()
	user.TokenStep = 0
	res := s.repo.Create(user)
	return res.Error
}

func (s *UserService) FindById(uid string) (database.User, error) {
	var user database.User
	res := s.repo.Where("uid = ?", uid).First(&user)
	if res.Error != nil {
		return database.User{}, errors.New("users not found")
	}
	return user, nil
}

func (s *UserService) DeleteById(uid string) error {
	res := s.repo.Delete(&database.User{}, "uid = ?", uid)
	return res.Error
}



