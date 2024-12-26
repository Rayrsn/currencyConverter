package currency

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func NewConverter() *Converter {
    return &Converter{
        BaseURL: "https://cdn.jsdelivr.net/gh/fawazahmed0/currency-api@1/latest/currencies/",
    }
}

func (c *Converter) Convert(amount float64, from, to string) (float64, error) {
    url := c.BaseURL + strings.ToLower(from) + "/" + strings.ToLower(to) + ".min.json"
    
    response, err := http.Get(url)
    if err != nil {
        return 0, err
    }
    defer response.Body.Close()

    var result map[string]interface{}
    if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
        return 0, err
    }

    rate, ok := result[strings.ToLower(to)].(float64)
    if !ok {
        return 0, fmt.Errorf("invalid rate conversion")
    }

    return amount * rate, nil
}
