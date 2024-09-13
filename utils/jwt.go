package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const ShortTokenSecretKey = "MainKhiladiTuAnadi"
const LongTokenSecretKey = "TheNewDays"

func GenerateShortToken(email string) (string, error) {

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 2).Unix(),
	})
	return jwtToken.SignedString([]byte(ShortTokenSecretKey))
}

func GenerateLongToken(email string) (string, error) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	})
	return jwtToken.SignedString([]byte(LongTokenSecretKey))
}

func VerifyShortToken(token string) (string, error) {
	parsedToken, err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		_, ok := jwtToken.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing")
		}
		return []byte(ShortTokenSecretKey), nil
	})
	if err != nil {
		return "", errors.New("could not parsed token")
	}
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return "", errors.New("token is not valid")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token")
	}
	email := claims["email"].(string)
	return email, nil
}

func VerifyLongToken(token string) (string, error){
	parsedToken,err := jwt.Parse(token, func(jwtToken *jwt.Token) (interface{}, error) {
		_,ok := jwtToken.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing")
		}
		return []byte(LongTokenSecretKey), nil
	})
	if err != nil {
		return "", errors.New("could not parsed token")
	}
	tokenIsValid := parsedToken.Valid
	if !tokenIsValid {
		return "", errors.New("token is not valid")
	}
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token")
	}
	email := claims["email"].(string)
	return email, nil
}