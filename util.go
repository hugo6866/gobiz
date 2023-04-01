package gobizApi

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var IsDebug bool = false

func marshalJSON(data interface{}) ([]byte, error) {
	jsonValue, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Error marshaling JSON data: %v", err)
	}
	return jsonValue, nil
}

func unmarshalJSON(data []byte, target interface{}) error {
	err := json.Unmarshal(data, target)
	if err != nil {
		return fmt.Errorf("Error unmarshaling JSON data: %v", err)
	}
	return nil
}

func printError(err error) {
	if IsDebug {
		if err != nil {
			fmt.Println("Error : ", err)
		}
	}
}
func printResponse(ok string) {
	if IsDebug {
		if ok != "" {
			fmt.Println("Response : ", ok)
		}
	}
}

func sendRequest(req *http.Request, authorization string) ([]byte, error) {
	setRequestHeaders(req, authorization, true)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		printResponse("Debug Response : " + string(body))
		return body, fmt.Errorf("Request failed with status: %d", resp.StatusCode)
	}

	return body, nil
}

func sendRequestGuest(req *http.Request) ([]byte, error) {
	setRequestHeaders(req, "", false)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("Error reading response: %v", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		printResponse("Debug Response: " + string(body))
		return body, fmt.Errorf("Request failed with status: %d", resp.StatusCode)
	}

	return body, nil
}
func setRequestHeaders(req *http.Request, authorization string, isAuthenticated bool) {
	req.Header.Set("X-User-Type", "merchant")
	req.Header.Set("X-Client-Id", "go-biz-mobile")
	req.Header.Set("X-Client-Secret", "sPC0qVk7gi76JUoGVfOfcgd7FfuaBv")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("User-Agent", "okhttp/3.12.10")

	if isAuthenticated {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authorization))
	} else {
		req.Header.Set("X-Platform", "Android")
		req.Header.Set("X-Appversion", "4.4.0")
		req.Header.Set("X-Appid", "com.gojek.resto")

	}
}
