package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/charmbracelet/huh"
)

var (
	baseCurrency string
	toCurrency   string
	amount       string
)

func main() {
	form := huh.NewForm(
		huh.NewGroup(
			// Ask the user for a base currency.
			huh.NewSelect[string]().
				Title("What is your base Currency").
				Options(
					huh.NewOption("$ USD", "USD"),
					huh.NewOption("£ GBP", "GBP"),
					huh.NewOption("€ EUR", "EUR"),
					huh.NewOption("¥ JPY", "JPY"),
				).
				Value(&baseCurrency),

			// Ask the user for a to currency.
			huh.NewSelect[string]().
				Title("What do you want to convert ino").
				Options(
					huh.NewOption("$ USD", "USD"),
					huh.NewOption("£ GBP", "GBP"),
					huh.NewOption("€ EUR", "EUR"),
					huh.NewOption("¥ JPY", "JPY"),
				).
				Value(&toCurrency),

			huh.NewInput().
				Title("How much to convert").
				Value(&amount).
				Validate(func(str string) error {
					_, err := strconv.Atoi(str)
					if err != nil {
						return errors.New("sorry, please enter a number")
					}
					return nil
				}),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	callExternalAPI(baseCurrency, toCurrency, amount)
}

type ExchangeResp struct {
	Disclaimer string
	License    string
	Timestamp  int
	Base       string
	Rates      map[string]float64
}

func callExternalAPI(baseCurrency string, toCurrency string, amount string) {
	apiURL := "https://openexchangerates.org/api/latest.json?app_id=de468c78c0ca4b4fa11a6457e4ebe7a0"

	er := new(ExchangeResp)
	getJSON(apiURL, er)

	rateFrom := er.Rates[baseCurrency]
	rateTo := er.Rates[toCurrency]

	a := convertToInt(amount)

	converted := a * (rateTo / rateFrom)

	fmt.Println(amount + " " + baseCurrency + " converted to " + toCurrency + " is " + strconv.FormatFloat(converted, 'f', -1, 64))
}

func getJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

func convertToInt(v string) float64 {
	i, err := strconv.ParseFloat(v, 64)
	if err != nil {
		log.Fatal(err)
	}

	return i
}
