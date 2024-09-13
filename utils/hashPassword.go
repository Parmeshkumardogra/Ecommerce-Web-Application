package utils

import (
	"errors"
	"log"

	"golang.org/x/crypto/bcrypt"
)
func HashPassword(password string) (string, error){
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password),10);
	if err != nil {
		log.Fatalf("unable to generate hashed password : %v",err.Error());
		return "", errors.New("unable to generate hashed password");
	};
	return string(hashedPassword),nil;
}

func CheckPassword(password, hashPassword string) (bool) {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword),[]byte(password));
	return err == nil;
}