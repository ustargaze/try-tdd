package main

import "testing"

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
