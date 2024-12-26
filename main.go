package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Rayrsn/currencyConverter/pkg/currency"
	"github.com/Rayrsn/currencyConverter/pkg/utils"
)

var Version = "1.1.0"

func main() {
    if err := run(); err != nil {
        fmt.Fprintf(os.Stderr, "error: %v\n", err)
        os.Exit(1)
    }
}

func run() error {
    if len(os.Args) == 1 {
        return fmt.Errorf("no arguments passed")
    }

    if len(os.Args) == 2 && (os.Args[1] == "-v" || os.Args[1] == "--version") {
        fmt.Println("Version:", Version)
        return nil
    }

    converter := currency.NewConverter()

    
	firstCurrency := os.Args[1]
	secondCurrency := os.Args[2]

    switch {
    case len(os.Args) == 3 && utils.IsCurrency(strings.ToUpper(firstCurrency)) && utils.IsCurrency(strings.ToUpper(secondCurrency)):
        rate, err := converter.Convert(1.0, firstCurrency, secondCurrency)
        if err != nil {
            return err
        }
        fmt.Printf("1 %s = %f %s\n", strings.ToUpper(firstCurrency), rate, strings.ToUpper(secondCurrency))
        
    case len(os.Args) == 4 && utils.IsNumber(firstCurrency):
		thirdCurrency := os.Args[3]
        amount, _ := strconv.ParseFloat(firstCurrency, 64)
        rate, err := converter.Convert(amount, secondCurrency, thirdCurrency)
        if err != nil {
            return err
        }
        fmt.Printf("%g %s = %f %s\n", amount, secondCurrency, rate, strings.ToUpper(thirdCurrency))
        
    default:
        return fmt.Errorf("invalid arguments")
    }

    return nil
}
