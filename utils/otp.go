package utils

import (
	"fmt"
	"math/rand"
)

func GenerateOTP() (string){
	randomOTP:=rand.Intn(1000000);
	otp:= fmt.Sprintf("%06d",randomOTP)
	fmt.Println(otp);
	return otp;
	
}