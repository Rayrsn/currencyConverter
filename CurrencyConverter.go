package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	if len(os.Args) > 4 {
		fmt.Println("Excess aguments passed")
		os.Exit(1)
	} else if len(os.Args) == 1 {
		fmt.Println("No arguments passed")
		os.Exit(1)
	} else if len(os.Args) == 2 {
		fmt.Println("Only one argument passed")
		os.Exit(1)
	}
	var mainUrl string = "https://api.exchangerate.host/latest"
	var response *http.Response
	var err error
	var url string

	if checkIfArgIsCurrency(getFirstArg()) {
		if checkIfArgIsCurrency(getSecondArg()) {
			url = mainUrl + "?" + "base=" + getFirstArg()
		}
	}

	// get the response from the link
	response, err = http.Get(url)
	if err != nil {
		fmt.Println(err)
	}

	// parse the response as json
	var exchangeRate map[string]interface{}
	json.NewDecoder(response.Body).Decode(&exchangeRate)
	// get the value of the key "rates"
	var rates map[string]interface{}
	rates = exchangeRate["rates"].(map[string]interface{})
	// print the values
	printResult(rates)
}

func getFirstArg() string {
	return os.Args[1]
}

func getSecondArg() string {
	return os.Args[2]
}

func printResult(rates map[string]interface{}) {
	if checkIfArgIsCurrency(getFirstArg()) {
		fmt.Println("1", getFirstArg(), "=", rates[getSecondArg()], getSecondArg())
	}
}

func checkIfArgIsCurrency(arg string) bool {
	switch arg {
	case "AED":
		return true

	case "AFN":
		return true

	case "ALL":
		return true

	case "AMD":
		return true

	case "ANG":
		return true

	case "AOA":
		return true

	case "ARS":
		return true

	case "AUD":
		return true

	case "AWG":
		return true

	case "AZN":
		return true

	case "BAM":
		return true

	case "BBD":
		return true

	case "BDT":
		return true

	case "BGN":
		return true

	case "BHD":
		return true

	case "BIF":
		return true

	case "BMD":
		return true

	case "BND":
		return true

	case "BOB":
		return true

	case "BRL":
		return true

	case "BSD":
		return true

	case "BTC":
		return true

	case "BTN":
		return true

	case "BWP":
		return true

	case "BYN":
		return true

	case "BZD":
		return true

	case "CAD":
		return true

	case "CDF":
		return true

	case "CHF":
		return true

	case "CLF":
		return true

	case "CLP":
		return true

	case "CNH":
		return true

	case "CNY":
		return true

	case "COP":
		return true

	case "CRC":
		return true

	case "CUC":
		return true

	case "CUP":
		return true

	case "CVE":
		return true

	case "CZK":
		return true

	case "DJF":
		return true

	case "DKK":
		return true

	case "DOP":
		return true

	case "DZD":
		return true

	case "EGP":
		return true

	case "ERN":
		return true

	case "ETB":
		return true

	case "EUR":
		return true

	case "FJD":
		return true

	case "FKP":
		return true

	case "GBP":
		return true

	case "GEL":
		return true

	case "GGP":
		return true

	case "GHS":
		return true

	case "GIP":
		return true

	case "GMD":
		return true

	case "GNF":
		return true

	case "GTQ":
		return true

	case "GYD":
		return true

	case "HKD":
		return true

	case "HNL":
		return true

	case "HRK":
		return true

	case "HTG":
		return true

	case "HUF":
		return true

	case "IDR":
		return true

	case "ILS":
		return true

	case "IMP":
		return true

	case "INR":
		return true

	case "IQD":
		return true

	case "IRR":
		return true

	case "ISK":
		return true

	case "JEP":
		return true

	case "JMD":
		return true

	case "JOD":
		return true

	case "JPY":
		return true

	case "KES":
		return true

	case "KGS":
		return true

	case "KHR":
		return true

	case "KMF":
		return true

	case "KPW":
		return true

	case "KRW":
		return true

	case "KWD":
		return true

	case "KYD":
		return true

	case "KZT":
		return true

	case "LAK":
		return true

	case "LBP":
		return true

	case "LKR":
		return true

	case "LRD":
		return true

	case "LSL":
		return true

	case "LYD":
		return true

	case "MAD":
		return true

	case "MDL":
		return true

	case "MGA":
		return true

	case "MKD":
		return true

	case "MMK":
		return true

	case "MNT":
		return true

	case "MOP":
		return true

	case "MRU":
		return true

	case "MUR":
		return true

	case "MVR":
		return true

	case "MWK":
		return true

	case "MXN":
		return true

	case "MYR":
		return true

	case "MZN":
		return true

	case "NAD":
		return true

	case "NGN":
		return true

	case "NIO":
		return true

	case "NOK":
		return true

	case "NPR":
		return true

	case "NZD":
		return true

	case "OMR":
		return true

	case "PAB":
		return true

	case "PEN":
		return true

	case "PGK":
		return true

	case "PHP":
		return true

	case "PKR":
		return true

	case "PLN":
		return true

	case "PYG":
		return true

	case "QAR":
		return true

	case "RON":
		return true

	case "RSD":
		return true

	case "RUB":
		return true

	case "RWF":
		return true

	case "SAR":
		return true

	case "SBD":
		return true

	case "SCR":
		return true

	case "SDG":
		return true

	case "SEK":
		return true

	case "SGD":
		return true

	case "SHP":
		return true

	case "SLL":
		return true

	case "SOS":
		return true

	case "SRD":
		return true

	case "SSP":
		return true

	case "STD":
		return true

	case "STN":
		return true

	case "SVC":
		return true

	case "SYP":
		return true

	case "SZL":
		return true

	case "THB":
		return true

	case "TJS":
		return true

	case "TMT":
		return true

	case "TND":
		return true

	case "TOP":
		return true

	case "TRY":
		return true

	case "TTD":
		return true

	case "TWD":
		return true

	case "TZS":
		return true

	case "UAH":
		return true

	case "UGX":
		return true

	case "USD":
		return true

	case "UYU":
		return true

	case "UZS":
		return true

	case "VES":
		return true

	case "VND":
		return true

	case "VUV":
		return true

	case "WST":
		return true

	case "XAF":
		return true

	case "XAG":
		return true

	case "XAU":
		return true

	case "XCD":
		return true

	case "XDR":
		return true

	case "XOF":
		return true

	case "XPD":
		return true

	case "XPF":
		return true

	case "XPT":
		return true

	case "YER":
		return true

	case "ZAR":
		return true

	case "ZMW":
		return true

	case "ZWL":
		return true

	default:
		return false
	}
}
