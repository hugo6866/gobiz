package gobizApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type GojekAuth struct {
	deviceOS   string
	phoneMake  string
	phoneModel string
	uniqueID   string
}

type OtpResponse struct {
	Data struct {
		OtpToken string `json:"otp_token"`
	} `json:"data"`
	Success bool     `json:"success"`
	Errors  []string `json:"errors"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
}

func NewGojekAuth(deviceOS, phoneMake, phoneModel, uniqueID string) *GojekAuth {
	return &GojekAuth{
		deviceOS:   deviceOS,
		phoneMake:  phoneMake,
		phoneModel: phoneModel,
		uniqueID:   uniqueID,
	}
}

func (g *GojekAuth) GetUniqueID() string {
	return g.uniqueID
}

func (g *GojekAuth) SetUniqueID(uniqueID string) {
	g.uniqueID = uniqueID
}

func (g *GojekAuth) setHeaders(req *http.Request) {
	req.Header.Set("X-Uniqueid", g.uniqueID)
	req.Header.Set("X-Deviceos", g.deviceOS)
	req.Header.Set("X-Phonemake", g.phoneMake)
	req.Header.Set("X-Phonemodel", g.phoneModel)
}

func (g *GojekAuth) GetOtpToken(number int) (string, error) {
	data := map[string]string{
		"phone_number":  fmt.Sprintf("%d", number),
		"client_secret": "sPC0qVk7gi76JUoGVfOfcgd7FfuaBv",
		"country_code":  "+62",
		"client_id":     "go-biz-mobile",
	}
	jsonValue, err := json.Marshal(data)
	if err != nil {
		printError(err)
		return "", err
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("https://goid.gojekapi.com/goid/login/request"), bytes.NewBuffer(jsonValue))
	if err != nil {
		printError(err)
		return "", err
	}

	g.setHeaders(req)
	body, err := sendRequestGuest(req)
	if err != nil {
		printResponse(string(body))
		printError(err)
		return "", err
	}
	var otpResponse OtpResponse
	err = json.Unmarshal(body, &otpResponse)
	if err != nil {
		printError(err)
		return "", err
	}
	printResponse(string(body))
	return otpResponse.Data.OtpToken, nil
}

func (g *GojekAuth) ProcessOtpToken(otp int, otptoken string) (string, error) {
	data := map[string]interface{}{
		"client_secret": "sPC0qVk7gi76JUoGVfOfcgd7FfuaBv",
		"grant_type":    "otp",
		"data": map[string]string{
			"otp_token": otptoken,
			"otp":       fmt.Sprintf("%d", otp),
		},
		"client_id": "go-biz-mobile",
	}

	jsonValue, err := json.Marshal(data)
	if err != nil {
		printError(err)
		return "", err
	}
	req, err := http.NewRequest("POST", fmt.Sprintf("https://goid.gojekapi.com/goid/token"), bytes.NewBuffer(jsonValue))
	if err != nil {
		printError(err)
		return "", err
	}

	g.setHeaders(req)
	body, err := sendRequestGuest(req)
	if err != nil {
		printError(err)
		return "", err
	}
	var tokenResponse TokenResponse
	err = json.Unmarshal(body, &tokenResponse)
	if err != nil {
		printError(err)
		return "", err
	}
	printResponse(string(body))
	return tokenResponse.AccessToken, nil
}
