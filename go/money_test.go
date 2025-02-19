package main

import (
	"testing"
	"try-tdd/stocks"
)

/**
function:
5 USD x 2 = 10 USD ✅
10 EUR x 2 = 20 EUR ✅
4002 KRW / 4 = 1000.5 KRW ✅
5 USD + 10 EUR = 17 USD
1 USD + 1100 KRW = 2000 KRW
*/

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

func assertEqual(t *testing.T, excepted stocks.Money, actual stocks.Money) {
	if excepted != actual {
		t.Errorf("Excepted %+v Got: %+v", excepted, actual)
	}
}

func TestAddition(t *testing.T) {
	var portfolio stocks.Portfolio
	var portfolioInDollars stocks.Money

	fiveDollars := stocks.NewMoney(5, "USD")
	tenDollars := stocks.NewMoney(10, "USD")
	fifteenDollars := stocks.NewMoney(15, "USD")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)
	portfolioInDollars = portfolio.Evaluate("USD")

	assertEqual(t, fifteenDollars, portfolioInDollars)
}
