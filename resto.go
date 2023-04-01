package gobizApi

import (
	"bytes"
	"fmt"
	"net/http"
)

type RestaurantStatusResponse struct {
	OpenStatus struct {
		Status    string `json:"status"`
		NextOpen  string `json:"next_open"`
		NextClose string `json:"next_close"`
	} `json:"open_status"`
	ForceClose bool `json:"force_close"`
}

func GetRestoStatus(restaurantID, authorization string) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.gojekapi.com/gofood/merchant/v2/restaurants/%s/open_status", restaurantID), nil)
	printError(err)

	body, err := sendRequest(req, authorization)
	if err != nil {
		printError(err)
		return "", err
	}
	printResponse(string(body))
	var statusResponse RestaurantStatusResponse
	err = unmarshalJSON(body, &statusResponse)
	if err != nil {
		printError(err)
		return "", err
	}
	return statusResponse.OpenStatus.Status, nil
}

func SetRestoClose(restaurantID, authorization string, status bool) (bool, error) {
	jsonData := map[string]bool{
		"force_close": status,
	}
	jsonValue, err := marshalJSON(jsonData)
	if err != nil {
		printError(err)
		return false, err
	}

	req, err := http.NewRequest("PATCH", fmt.Sprintf("https://api.gojekapi.com/gofood/merchant/v2/restaurants/%s", restaurantID), bytes.NewBuffer(jsonValue))
	if err != nil {
		printError(err)
		return false, err
	}

	body, err := sendRequest(req, authorization)
	if err != nil {
		printError(err)
		return false, err
	}

	var statusResponse RestaurantStatusResponse
	err = unmarshalJSON(body, &statusResponse)
	if err != nil {
		printError(err)
		return false, err
	}

	return statusResponse.ForceClose, nil
}
