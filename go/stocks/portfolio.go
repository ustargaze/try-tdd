package stocks

import "errors"

type Portfolio []Money

func (p Portfolio) Add(money Money) Portfolio {
	return append(p, money)
}

func (p Portfolio) Evaluate(currency string) (Money, error) {
	total := 0.0
	failedConversions := make([]string, 0)
	for _, money := range p {
		if convertedAmount, ok := convert(money, currency); ok {
			total += convertedAmount
		} else {
			failedConversions = append(failedConversions, money.currency+"->"+currency)
		}
	}
	if len(failedConversions) == 0 {
		return NewMoney(total, currency), nil
	}
	failures := "["
	for _, f := range failedConversions {
		failures += f + ","
	}
	failures += "]"
	return NewMoney(0, ""), errors.New("Missing exchange rate(s):" + failures)
}

func convert(money Money, currency string) (float64, bool) {
	exchangeRate := map[string]float64{
		"EUR->USD": 1.2,
		"USD->KRW": 1100,
	}
	if money.currency == currency {
		return money.amount, true
	}
	key := money.currency + "->" + currency
	rate, ok := exchangeRate[key]
	return money.amount * rate, ok
}
