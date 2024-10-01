package genUnqAccNo

import (
	"strconv"
	"time"
	"fmt"
)

func GenerateUniqueAccountNo(accountType string) string{

	now := time.Now();
	timePart :=  now.Format("20060102150405");
	nanoPart := strconv.Itoa(now.Nanosecond())[0:4];
	var accountCode string;
	switch accountType {
	case "saving":
		accountCode = "SV";
	case "credit":
		accountCode = "CR";
	case "current":
		accountCode = "CU";
	case "fixed":
		accountCode = "FD";
	default:
		return accountCode;
	}
	accountNumber := accountCode + timePart + nanoPart;
	fmt.Println("Unique Account Number",accountNumber);
	return accountNumber

}
