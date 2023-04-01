package gobizApi

import (
	"bytes"
	"net/http"
	"time"
)

func GetRestoTransactions(restaurantID, authorization string, start, end string) (string, error) {
	startStr := start + "T00:00:00-05:00"
	endStr := end + "T23:59:59-05:00"
	jsonData := `{"from":0,"included_categories":{"incoming":["transaction_share","action"]},"query":[{"clauses":[{"field":"metadata.transaction.status","op":"in","value":["settlement","capture","cancel","failure","deny"]},{"field":"metadata.transaction.merchant_id","op":"equal","value":"G194812516"}],"op":"and"}],"size":10,"sort":{"time":{"order":"desc"}},"source":["metadata.source","metadata.transaction.id","metadata.transaction.transaction_source.source","metadata.transaction.transaction_source.service","metadata.transaction.order_id","metadata.transaction.transaction_time","metadata.transaction.real_gross_amount","metadata.transaction.payment_type","metadata.transaction.pop.id","metadata.card_type","metadata.transaction_type","metadata.transaction.status","metadata.merchant_mid","metadata.merchant_tid","metadata.merchant_code","metadata.batch_number","metadata.reference_number","metadata.reference_id","metadata.auth_id_response","metadata.transaction.status","metadata.transaction_number","metadata.masked_card","metadata.transaction.customer","metadata.transaction.custom_field1","metadata.transaction.custom_field2","metadata.transaction.custom_field3","metadata.transaction.metadata","metadata.acquiring_bank","metadata.on_us"],"time_range":{"gte":"` + startStr + `","lte":"` + endStr + `"}}`
	jsonValue := []byte(jsonData)
	req, err := http.NewRequest("POST", "https://api.gobiz.co.id/journals/search", bytes.NewBuffer(jsonValue))
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

func GetTodayTransactions(restaurantID, authorization string) {
	now := time.Now()
	start := now.Format("2006-01-02")
	end := now.Format("2006-01-02")
	GetRestoTransactions(restaurantID, authorization, start, end)
}
