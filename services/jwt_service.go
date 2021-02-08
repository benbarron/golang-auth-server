package services

import (
	"errors"
	"github.com/benbarron/UserMicroService/database"
	"github.com/dgrijalva/jwt-go"
	"os"
	"strconv"
	"time"
)

type JwtService struct {
	SecretKey       		string
	Issuer          		string
	AccessExpirationHours 	int64
	RefreshExpirationHours 	int64
}

type JwtClaim struct {
	User database.User
	jwt.StandardClaims
}

func NewJwtService() *JwtService {
	refreshExp, _ := strconv.ParseInt(os.Getenv("JWT_REFRESH_EXP"), 10, 32)
	accessExp, _ := strconv.ParseInt(os.Getenv("JWT_ACCESS_EXP"), 10, 32)
	issuer := os.Getenv("JWT_ISSUER")
	secretKey := os.Getenv("JWT_SECRET_KEY")

	return &JwtService{
		SecretKey: secretKey,
		Issuer: issuer,
		AccessExpirationHours: accessExp,
		RefreshExpirationHours: refreshExp,
	}
}

func (j *JwtService) GenerateRefreshToken(user database.User) (signedToken string, err error) {
	claims := &JwtClaim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.RefreshExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(j.SecretKey))

	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (j *JwtService) GenerateAccessToken(user database.User) (signedToken string, err error) {
	claims := &JwtClaim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(j.AccessExpirationHours)).Unix(),
			Issuer:    j.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(j.SecretKey))

	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (j *JwtService) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)

	if err != nil {
		return
	}

	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("Couldn't parse claims")
		return
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("JWT is expired")
		return
	}

	return

}