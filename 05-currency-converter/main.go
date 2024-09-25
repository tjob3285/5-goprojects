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

/*var (
	burger       string
	toppings     []string
	sauceLevel   int
	name         string
	instructions string
	discount     bool
)

func main() {
	form := huh.NewForm(
		huh.NewGroup(
			// Ask the user for a base burger and toppings.
			huh.NewSelect[string]().
				Title("Choose your burger").
				Options(
					huh.NewOption("Charmburger Classic", "classic"),
					huh.NewOption("Chickwich", "chickwich"),
					huh.NewOption("Fishburger", "fishburger"),
					huh.NewOption("Charmpossible™ Burger", "charmpossible"),
				).
				Value(&burger), // store the chosen option in the "burger" variable

			// Let the user select multiple toppings.
			huh.NewMultiSelect[string]().
				Title("Toppings").
				Options(
					huh.NewOption("Lettuce", "lettuce").Selected(true),
					huh.NewOption("Tomatoes", "tomatoes").Selected(true),
					huh.NewOption("Jalapeños", "jalapeños"),
					huh.NewOption("Cheese", "cheese"),
					huh.NewOption("Vegan Cheese", "vegan cheese"),
					huh.NewOption("Nutella", "nutella"),
				).
				Limit(4). // there’s a 4 topping limit!
				Value(&toppings),

			// Option values in selects and multi selects can be any type you
			// want. We’ve been recording strings above, but here we’ll store
			// answers as integers. Note the generic "[int]" directive below.
			huh.NewSelect[int]().
				Title("How much Charm Sauce do you want?").
				Options(
					huh.NewOption("None", 0),
					huh.NewOption("A little", 1),
					huh.NewOption("A lot", 2),
				).
				Value(&sauceLevel),
		),

		// Gather some final details about the order.
		huh.NewGroup(
			huh.NewInput().
				Title("What’s your name?").
				Value(&name).
				// Validating fields is easy. The form will mark erroneous fields
				// and display error messages accordingly.
				Validate(func(str string) error {
					if str == "Frank" {
						return errors.New("sorry, we don’t serve customers named Frank")
					}
					return nil
				}),

			huh.NewText().
				Title("Special Instructions").
				CharLimit(400).
				Value(&instructions),

			huh.NewConfirm().
				Title("Would you like 15% off?").
				Value(&discount),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}

	if !discount {
		fmt.Println("What? You didn’t take the discount?!")
	}
}*/

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
	Disclaimer string             `json:"disclaimer"`
	License    string             `json:"license"`
	Timestamp  int                `json:"timestamp"`
	Base       string             `json:"base"`
	Rates      map[string]float64 `json:"map[string]float64"`
}

func callExternalAPI(baseCurrency string, toCurrency string, amount string) {
	//apiURL := "https://openexchangerates.org/api/convert/" + amount + "/" + baseCurrency + "/" + toCurrency + "?app_id=de468c78c0ca4b4fa11a6457e4ebe7a0"
	apiURL := "https://openexchangerates.org/api/latest.json?app_id=de468c78c0ca4b4fa11a6457e4ebe7a0"

	er := new(ExchangeResp)
	getJSON(apiURL, er)
	fmt.Println(er.Rates[baseCurrency])
	fmt.Println(er.Rates[toCurrency])
	bc := strconv.Itoa(int(er.Rates[baseCurrency]))
	tc := strconv.Itoa(int(er.Rates[toCurrency]))

	fmt.Print(bc + " converts to " + amount + " in " + tc)
}

func getJSON(url string, target interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
