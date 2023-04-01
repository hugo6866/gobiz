package gobizApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetRestoMenu(authorization, merchantID string) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.gojekapi.com/mpp/merchant/restaurant_menus?merchant_ids=%s", merchantID), nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return "", err

	}

	body, err := sendRequest(req, authorization)
	if err != nil {
		printError(err)
		return "", err
	}

	printResponse(string(body))
	return string(body), nil
}

func SetStockAvailable(authorization, itemID string, isAvailable bool) (bool, error) {
	jsonData := map[string]bool{
		"in_stock": isAvailable,
	}
	jsonValue, err := json.Marshal(jsonData)
	req, err := http.NewRequest("PATCH", fmt.Sprintf("https://api.gojekapi.com/gofood/merchant/v1/menu_items/%s", itemID), bytes.NewBuffer(jsonValue))
	fmt.Printf("https://api.gojekapi.com/gofood/merchant/v1/menu_items/%s\n", itemID)
	if err != nil {
		printError(err)
		return false, err
	}

	body, err := sendRequest(req, authorization)
	if err != nil {
		printError(err)
		return false, err
	}

	printResponse(string(body))
	return true, nil
}
