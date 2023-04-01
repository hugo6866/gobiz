package main

import (
	"fmt"

	gobizApi "github.com/hugo6866/gobiz"
)

func main() {
	// Create an instance of the GojekAPI class with a unique ID
	api := gobizApi.NewGojekAuth("Android 11", "Samsung", "Galaxy S22", "1453455d71e619")
	var phoneNumber int
	fmt.Print("Enter your phone number: +62")
	fmt.Scan(&phoneNumber)

	// Get an OTP token
	otpToken, err := api.GetOtpToken(phoneNumber)
	if err != nil {
		fmt.Printf("Error getting OTP token: %v\n", err)
		return
	} else {
		fmt.Println(otpToken)
	}

	// Input the OTP code
	var otpCode int
	fmt.Print("Enter your OTP code: ")
	fmt.Scan(&otpCode)

	// Use the OTP token to get the final authentication token
	authToken, err := api.ProcessOtpToken(otpCode, otpToken)
	if err != nil {
		fmt.Printf("Error getting authentication token: %v\n", err)
		return
	} else {
		fmt.Println(authToken)
	}
}
