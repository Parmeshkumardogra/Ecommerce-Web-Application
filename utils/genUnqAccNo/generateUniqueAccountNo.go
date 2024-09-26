package genUnqAccNo

import (
	"strconv"
	"time"
	"fmt"
)

func GenerateUniqueAccountNo(accountType string) string{

	now := time.Now();
	timePart :=  now.Format("20060102150405");
	nanoPart := strconv.Itoa(now.Nanosecond())[0:3];
	var accountCode string;
	switch accountType {
	case "saving":
		accountCode = "SAV";
	case "credit":
		accountCode = "CRE";
	case "current":
		accountCode = "CUR";
	case "fixed":
		accountCode = "FIX";
	default:
		return accountCode;
	}
	accountNumber := accountCode + timePart + nanoPart;
	fmt.Println("Unique Account Number",accountNumber);
	return accountNumber

}
