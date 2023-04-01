package main

import (
	"fmt"

	gobizApi "github.com/hugo6866/gobiz"
)

func main() {
	// Set the authorization key
	authorization := "YOUR_AUTHORIZATION_KEY"

	// Get the list of merchants
	merchants, err := gobizApi.GetMerchants(authorization)
	if err != nil {
		fmt.Printf("Error retrieving list of merchants: %v\n", err)
		return
	} else {
		fmt.Println(merchants)
	}
}
