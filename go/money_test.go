package main

import (
	"reflect"
	"testing"
	"try-tdd/stocks"
)

var bank stocks.Bank

func initExchangeRates() {
	bank = stocks.NewBank()
	bank.AddExchangeRate("EUR", "USD", 1.2)
	bank.AddExchangeRate("USD", "KRW", 1100)
}

func TestMultiplicationDollars(t *testing.T) {
	money := stocks.NewMoney(5, "USD")
	actual := money.Times(2)
	excepted := stocks.NewMoney(10, "USD")

	assertEqual(t, excepted, actual)
}

func TestMultiplicationInEuros(t *testing.T) {
	tenEuros := stocks.NewMoney(10, "EUR")
	actual := tenEuros.Times(2)
	excepted := stocks.NewMoney(20, "EUR")

	assertEqual(t, excepted, actual)
}

func TestDivision(t *testing.T) {
	money := stocks.NewMoney(4002, "KRW")
	actual := money.Divide(4)
	excepted := stocks.NewMoney(1000.5, "KRW")

	assertEqual(t, excepted, actual)
}

func assertEqual(t *testing.T, excepted, actual interface{}) {
	if excepted != actual {
		t.Errorf("Excepted %+v Got: %+v", excepted, actual)
	}
}

func TestAddition(t *testing.T) {
	initExchangeRates()
	var portfolio stocks.Portfolio

	fiveDollars := stocks.NewMoney(5, "USD")
	tenDollars := stocks.NewMoney(10, "USD")
	fifteenDollars := stocks.NewMoney(15, "USD")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)
	portfolioInDollars, err := portfolio.Evaluate(bank, "USD")

	assertNil(t, err)
	assertEqual(t, fifteenDollars, *portfolioInDollars)
}

func TestAdditionOfDollarsAndEuros(t *testing.T) {
	initExchangeRates()
	var portfolio stocks.Portfolio

	fiveDollars := stocks.NewMoney(5, "USD")
	tenEuros := stocks.NewMoney(10, "EUR")
	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenEuros)

	excepted := stocks.NewMoney(17, "USD")
	actual, err := portfolio.Evaluate(bank, "USD")

	assertNil(t, err)
	assertEqual(t, excepted, *actual)
}

func TestAdditionOfDollarsAndWons(t *testing.T) {
	initExchangeRates()
	var portfolio stocks.Portfolio

	oneDollar := stocks.NewMoney(1, "USD")
	elevenHundredWon := stocks.NewMoney(1100, "KRW")
	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(elevenHundredWon)

	excepted := stocks.NewMoney(2200, "KRW")
	actual, err := portfolio.Evaluate(bank, "KRW")

	assertNil(t, err)
	assertEqual(t, excepted, *actual)
}

func TestAdditionWithMultipleMissingExchangeRates(t *testing.T) {
	initExchangeRates()
	var portfolio stocks.Portfolio

	oneDollar := stocks.NewMoney(1, "USD")
	oneEur := stocks.NewMoney(1, "EUR")
	oneWon := stocks.NewMoney(1, "KRW")

	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(oneEur)
	portfolio = portfolio.Add(oneWon)

	exceptedErrorMessage := "Missing exchange rate(s):[USD->Kalganid,EUR->Kalganid,KRW->Kalganid,]"
	_, actualError := portfolio.Evaluate(bank, "Kalganid")

	if exceptedErrorMessage != actualError.Error() {
		t.Errorf("Excepted %s Got %s", exceptedErrorMessage, actualError.Error())
	}
}

func TestConversionWithDifferentRatesBetweenTwoCurrencies(t *testing.T) {
	initExchangeRates()
	tenEuros := stocks.NewMoney(10, "EUR")
	actual, err := bank.Convert(tenEuros, "USD")
	assertNil(t, err)
	assertEqual(t, stocks.NewMoney(12, "USD"), *actual)
	bank.AddExchangeRate("EUR", "USD", 1.3)
	actual, err = bank.Convert(tenEuros, "USD")
	assertNil(t, err)
	assertEqual(t, stocks.NewMoney(13, "USD"), *actual)
}

func TestWhatIsTheConversionRateFromEURToUSD(t *testing.T) {
	initExchangeRates()
	tenEuros := stocks.NewMoney(10, "EUR")
	actual, err := bank.Convert(tenEuros, "USD")
	assertNil(t, err)
	assertEqual(t, stocks.NewMoney(12, "USD"), *actual)
}

func assertNil(t *testing.T, actual interface{}) {
	if actual != nil && !reflect.ValueOf(actual).IsNil() {
		t.Errorf("Excepted to be nil, found: [%+v]", actual)
	}
}

func TestConversionWithMissingExchangeRate(t *testing.T) {
	initExchangeRates()
	tenEuros := stocks.NewMoney(10, "EUR")
	actual, err := bank.Convert(tenEuros, "Kalganid")
	assertNil(t, actual)
	assertEqual(t, "EUR->Kalganid", err.Error())
}
