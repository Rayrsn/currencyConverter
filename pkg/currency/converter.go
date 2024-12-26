package currency

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type CurrencyResponse struct {
    Date     string                 `json:"date"`
    Currency map[string]interface{} `json:",inline"`
}

func NewConverter() *Converter {
    return &Converter{
        BaseURL: "https://cdn.jsdelivr.net/gh/fawazahmed0/currency-api@1/latest/currencies/",
    }
}

func (c *Converter) Convert(amount float64, from, to string) (float64, error) {
    url := c.BaseURL + strings.ToLower(from) + ".min.json"
    
    response, err := http.Get(url)
    if err != nil {
        return 0, err
    }
    defer response.Body.Close()

    var result map[string]interface{}
    if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
        return 0, err
    }

    rates, exists := result[strings.ToLower(from)].(map[string]interface{})
    if !exists {
        return 0, fmt.Errorf("currency %s not found", from)
    }

    rate, exists := rates[strings.ToLower(to)].(float64)
    if !exists {
        return 0, fmt.Errorf("exchange rate for %s not found", to)
    }

    return amount * rate, nil
}
