# GoBiz Resto App Unofficial API Wrapper

This is an unofficial API wrapper for the [GoBiz Resto App](https://www.gobiz.co.id/) written in Golang.

## Installation

You can install the package using the `go get` command:
```
go get -u github.com/hugo6866/gobiz
```

## Usage

Here's an example code snippet that demonstrates how to use the wrapper :

## Example Usage: Getting an Authentication Token

The following example demonstrates how to use the `GojekAuth` struct to get an authentication token for the GoBiz REST API.

```go
import (
	"fmt"
	"github.com/hugo6866/gobiz"
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
```
More examples on how to use the GoBiz Resto App Unofficial API Wrapper, please check out the [examples folder](https://github.com/hugo6866/gobiz/tree/master/example) in the GitHub repository. 

## API Endpoints

The following API endpoints are currently supported by the wrapper :

### Gojek Authentication
- getOtpToken(number)
  - `number`: Phone number to request an OTP token
  - Returns: `otp_token` (string) and an error (if any)
- processOtpToken(otp, otpToken)
  - `otp`: The OTP received on the phone number
  - `otpToken`: The OTP token obtained from the `getOtpToken` method
  - Returns: `access_token` (string) and an error (if any)

### Orders
- fetchCompletedOrders(restaurantID, authorization)
  - `restaurantID`: Restaurant ID to get completed orders
  - `authorization`: API key for the restaurant
- fetchNewOrders(restaurantID, authorization)
  - `restaurantID`: Restaurant ID to get new orders
  - `authorization`: API key for the restaurant
- fetchCanceledOrders(restaurantID, authorization)
  - `restaurantID`: Restaurant ID to get canceled orders
  - `authorization`: API key for the restaurant

### Restaurant
- setRestoClose(restaurantID, authorization, isClose)
  - `restaurantID`: Restaurant ID to set close status
  - `authorization`: API key for the restaurant
  - `isClose`: `true` for closed, `false` for open
- getRestoStatus(restaurantID, authorization)
  - `restaurantID`: Restaurant ID to get status
  - `authorization`: API key for the restaurant
- getRestoTransactions(restaurantID, authorization, start, end)
  - `restaurantID`: Restaurant ID to get transactions
  - `authorization`: API key for the restaurant
  - `start`: Start date in "yyyy-mm-dd" format
  - `end`: End date in "yyyy-mm-dd" format
- getTodayTransactions(restaurantID, authorization)
  - `restaurantID`: Restaurant ID to get today's transactions
  - `authorization`: API key for the restaurant

### Merchant
- getMerchants(authorization)
  - `authorization`: API key for the merchant
- getMerchantInfo(authorization, merchantID)
  - `authorization`: API key for the merchant
  - `merchantID`: Merchant ID to get information

### Menu
- getRestoMenu(authorization, restaurantID)
  - `authorization`: API key for the restaurant
  - `restaurantID`: Restaurant ID to get menu
- setStockAvailable(authorization, itemID, isAvailable)
  - `authorization`: API key for the restaurant
  - `itemID`: Item ID to set availability
  - `isAvailable`: `true` for available, `false` for not available

### Order Management
- acceptNewOrders(authorization, orderID)
  - `authorization`: API key for the restaurant
  - `orderID`: Order ID to accept
- setOrderPrepared(authorization, orderID)
  - `authorization`: API key for the restaurant
  - `orderID`: Order ID to set as prepared
- cancelOrder(authorization, orderID, reason, itemID, uuid)
  - `authorization`: API key for the restaurant
  - `orderID`: Order ID to cancel
  - `reason`: Cancellation reason
  - `itemID`: Item ID for "ITEMS_OUT_OF_STOCK" reason
  - `uuid`: UUID for "ITEMS_OUT_OF_STOCK" reason
- cancelOrderStockEmpty(authorization, orderID, itemID, uuid)
  - `authorization`: API key for the restaurant
  - `orderID`: Order ID to cancel due to out of stock
  - `itemID`: Out of stock item ID
  - `uuid`: Out of stock item UUID
- cancelOrderRestoClosed(authorization, orderID)
  - `authorization`: API key for the restaurant
  - `orderID`: Order ID to cancel due to closed restaurant
- cancelOrderHighDemand(authorization, orderID)
  - `authorization`: API key for the restaurant
  - `orderID`: Order ID to cancel due to high demand

### Debug Mode
- IsDebug: A global boolean variable to enable or disable debug mode.
  - `true`: Debug mode enabled
  - `false`: Debug mode disabled (default)

### Disclaimer
Please note that this is an unofficial API wrapper for the GoBiz Resto App and is not affiliated with or endorsed by GoBiz or its parent company, Gojek
