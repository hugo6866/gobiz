package gobizApi

import (
	"fmt"
	"net/http"
)

func GetMerchants(authorization string) (string, error) {
	req, err := http.NewRequest("GET", "https://api.gobiz.co.id/goresto/v5/public/users/config", nil)
	if err != nil {
		printError(fmt.Errorf("Error creating request: %v", err))
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

func GetMerchantInfo(authorization, merchantID string) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.gobiz.co.id/goauth/horme/merchants/%s", merchantID), nil)
	if err != nil {
		printError(fmt.Errorf("Error creating request: %v", err))
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
