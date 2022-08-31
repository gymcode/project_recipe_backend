package utils

import (
	"log"
	"strings"
)

func CountryValidation(msisdn, isoCode string) string {
	// switching cases for the isocodes
	var number string
	switch isoCode {
	// code for Ghana
	case "GH":
		var countryCode string = "233"
		number = msisdnValidation(msisdn, countryCode)
	}
	log.Println(number)
	return number
}

func msisdnValidation(msisdn, countryCode string) string {
	var mobileNumber string

	switch {

	case strings.HasPrefix(msisdn, "0") && len(msisdn) == 10:
		mobileNumber = countryCode + strings.TrimPrefix(msisdn, "0")
		//do something
	case strings.HasPrefix(msisdn, countryCode) && len(msisdn) == 12:
		mobileNumber = msisdn
	case strings.HasPrefix(msisdn, "+"+countryCode) && len(msisdn) > 12:
		mobileNumber = strings.TrimPrefix(msisdn, "+")
		// do something
	case strings.HasPrefix(msisdn, "0") == false && len(msisdn) == 9:
		mobileNumber = countryCode + msisdn
	}
	log.Println(mobileNumber)
	return mobileNumber
}
