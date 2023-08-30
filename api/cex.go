package api

import (
	"encoding/json"
	"fmt"
	"io"
	"jhonidev/go/goCryptocurrencies/datatypes"
	"net/http"
	"strings"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*datatypes.Rate, error) {
	if len(currency) < 3 {
		return nil, fmt.Errorf("Invalid currency value. Please check the currency value and its length. (%v received)", currency)
	}
	upperCurrency := strings.ToUpper(currency)
	res, err := http.Get(fmt.Sprintf(apiUrl, upperCurrency))

	if err != nil {
		return nil, err
	}

	var response CEXResponse
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		errJson := json.Unmarshal(bodyBytes, &response)
		if errJson != nil {
			return nil, errJson
		}
	} else {
		return nil, fmt.Errorf("Status code received: %v", res.StatusCode)
	}

	rate := datatypes.Rate{Currency: currency, Price: response.Bid}
	return &rate, nil

}
