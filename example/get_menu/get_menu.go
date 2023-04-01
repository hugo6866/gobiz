package main

import (
	"fmt"

	gobizApi "github.com/hugo6866/gobiz"
)

func main() {
	// Set the authorization key and merchant ID
	authorization := "YOUR_AUTHORIZATION_KEY"
	merchantID := "YOUR_MERCHANT_ID"

	// Get the restaurant's menu
	menu, err := gobizApi.GetRestoMenu(authorization, merchantID)
	if err != nil {
		fmt.Printf("Error retrieving restaurant menu: %v\n", err)
		return
	} else {
		fmt.Println(menu)
	}
}
