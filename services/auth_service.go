package services

import (
	"errors"
	"github.com/benbarron/golang-auth-server/database"
	"gorm.io/gorm"
)

type AuthService struct {
	AuthRepository *gorm.DB
	LoggingServicve *LoggingService
	JwtService *JwtService
}

func NewAuthService() *AuthService {
	return &AuthService{
		LoggingServicve: NewLoggingService("AuthService"),
		AuthRepository:   database.GetDatabase(),
		JwtService: NewJwtService(),
	}
}

func (a *AuthService) Login(username string, password string) (database.User, error) {
	var user database.User
	res := a.AuthRepository.Where("username = ?", username).First(&user)
	if res.Error != nil {
		return user, res.Error
	}
	if !user.CheckPassword(password) {
		return user, errors.New("invalid credentials")
	}
	return user, nil
}

func (a *AuthService) GetUserFromToken(token string) (database.User, error) {
	claims, err := a.JwtService.ValidateToken(token)
	if err != nil {
		return database.User{}, err
	}
	return claims.User, nil
}


