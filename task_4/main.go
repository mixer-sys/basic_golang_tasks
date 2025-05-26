package main

import "fmt"

type CurrencyConverter struct {
	Rate float64
}

func (cc *CurrencyConverter) ConvertToUSD(rubles float64) float64 {
	return rubles / cc.Rate
}

func (cc *CurrencyConverter) ConvertToRUB(dollars float64) float64 {
	return dollars * cc.Rate
}

func main() {
	converter := CurrencyConverter{Rate: 75.5}
	fmt.Printf("%.1f\n", converter.ConvertToUSD(755))
	fmt.Printf("%.1f\n", converter.ConvertToRUB(1000))
}
