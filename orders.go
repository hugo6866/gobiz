package gobizApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func FetchOrders(restaurantID, authorization, status string) (string, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://api.gojekapi.com/mocha/v3/orders?restaurant_id=%s&status=%s&limit=100&page=1", restaurantID, status), nil)
	if err != nil {
		printError(err)
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

func FetchNewOrders(restaurantID, authorization string) (string, error) {
	return FetchOrders(restaurantID, authorization, "AWAITING_ACCEPTANCE")
}

func FetchCompletedOrders(restaurantID, authorization string) (string, error) {
	return FetchOrders(restaurantID, authorization, "COMPLETED")
}

func FetchCanceledOrders(restaurantID, authorization string) (string, error) {
	return FetchOrders(restaurantID, authorization, "UNFULFILLED")
}

func CancelOrder(authorization, orderID, cancelReasonCode, id, uuid string) (bool, error) {
	data := map[string]interface{}{
		"cancel_reason_code": cancelReasonCode,
	}

	if cancelReasonCode == "ITEMS_OUT_OF_STOCK" {
		data["out_of_stock_items"] = []map[string]interface{}{
			{
				"id":   id,
				"uuid": uuid,
			},
		}
		data["out_of_stock_variants"] = []interface{}{}
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		printError(err)
		return false, err
	}

	req, err := http.NewRequest("PUT", fmt.Sprintf("https://api.gojekapi.com/buffet/v1/orders/%s/merchant/cancelled", orderID), bytes.NewBuffer(jsonData))
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

func CancelOrderStockEmpty(authorization, orderID, id, uuid string) (bool, error) {
	return CancelOrder(authorization, orderID, "ITEMS_OUT_OF_STOCK", id, uuid)
}

func CancelOrderRestoClosed(authorization, orderID string) (bool, error) {
	return CancelOrder(authorization, orderID, "RESTAURANT_CLOSED", "", "")
}

func CancelOrderHighDemand(authorization, orderID string) (bool, error) {
	return CancelOrder(authorization, orderID, "HIGH_DEMAND", "", "")
}
func SetOrderPrepared(authorization, orderID string) (string, error) {
	now := time.Now()

	data := map[string]interface{}{
		"timestamp":   now,
		"action_name": "FOOD_PREPARED_FOR_PICKUP",
	}

	jsonData, err := marshalJSON(data)
	if err != nil {
		printError(err)
		return "", err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("https://api.gojekapi.com/buffet/v1/orders/%s/merchant/food-prepared", orderID), bytes.NewBuffer(jsonData))
	if err != nil {
		printError(err)
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

func AcceptNewOrders(authorization, orderID string) (bool, error) {
	req, err := http.NewRequest("PUT", fmt.Sprintf("https://api.gojekapi.com/buffet/v1/orders/%s/merchant/accepted", orderID), nil)
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
