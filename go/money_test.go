package main

import "testing"

/**
function:
5 USD x 2 = 10 USD ✅
10 EUR x 2 = 20 EUR ✅
4002 KRW / 4 = 1000.5 KRW ✅
5 USD + 10 EUR = 17 USD
1 USD + 1100 KRW = 2000 KRW
*/

type Money struct {
	amount   float64
	currency string
}

func (m Money) Times(multiplier int) Money {
	return Money{amount: m.amount * float64(multiplier), currency: m.currency}
}

func (m Money) Divide(divisor int) Money {
	return Money{amount: m.amount / float64(divisor), currency: m.currency}
}

type Portfolio []Money

func (p Portfolio) Add(money Money) Portfolio {
	return append(p, money)
}

func (p Portfolio) Evaluate(currency string) Money {
	total := 0.0
	for _, money := range p {
		total += money.amount
	}
	return Money{amount: total, currency: currency}
}

func TestMultiplicationDollars(t *testing.T) {
	money := Money{amount: 5, currency: "USD"}
	actual := money.Times(2)
	excepted := Money{amount: 10, currency: "USD"}

	assertEqual(t, excepted, actual)
}

func TestMultiplicationInEuros(t *testing.T) {
	tenEuros := Money{amount: 10, currency: "EUR"}
	actual := tenEuros.Times(2)
	excepted := Money{amount: 20, currency: "EUR"}

	assertEqual(t, excepted, actual)
}

func TestDivision(t *testing.T) {
	money := Money{amount: 4002, currency: "KRW"}
	actual := money.Divide(4)
	excepted := Money{amount: 1000.5, currency: "KRW"}

	assertEqual(t, excepted, actual)
}

func assertEqual(t *testing.T, excepted Money, actual Money) {
	if excepted != actual {
		t.Errorf("Excepted %+v Got: %+v", excepted, actual)
	}
}

func TestAddition(t *testing.T) {
	var portfolio Portfolio
	var portfolioInDollars Money

	fiveDollars := Money{amount: 5, currency: "USD"}
	tenDollars := Money{amount: 10, currency: "USD"}
	fifteenDollars := Money{amount: 15, currency: "USD"}

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)
	portfolioInDollars = portfolio.Evaluate("USD")

	assertEqual(t, fifteenDollars, portfolioInDollars)
}
