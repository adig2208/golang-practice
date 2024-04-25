package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	revenue, err := getUserInput("Enter Revenue: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	taxRate, err := getUserInput("Enter Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}
	expenses, err := getUserInput("Enter Expenses: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, eat, ratio := calculateProfit(taxRate, revenue, expenses)

	fmt.Printf("Earnings before tax: %.1f \n", ebt)
	fmt.Printf("Earnings after tax (Profit): %.1f \n", eat)
	fmt.Printf("Ratio: %.3f ", ratio)
	storeResults(ebt, eat, ratio)
}

func getUserInput(text string) (float64, error) {
	var input float64
	fmt.Print(text)
	fmt.Scan(&input)
	if input <= 0 {
		return 0, errors.New("Invalid input. Value should be greater than 0.")
	}
	return input, nil
}

func storeResults(ebt, eat, ratio float64) {
	results := fmt.Sprintf("Earnings before tax: %.1f \nEarnings after tax (Profit): %.1f \nRatio: %.3f", ebt, eat, ratio)
	os.WriteFile("profits.txt", []byte(results), 0644)
}

func calculateProfit(taxRate, revenue, expenses float64) (float64, float64, float64) {
	var earningsBeforeTax float64 = revenue - expenses
	var taxAmount float64 = earningsBeforeTax * (1 - taxRate/100)
	var earningsAfterTax float64 = earningsBeforeTax - taxAmount
	var ratio float64 = earningsBeforeTax / earningsAfterTax
	return earningsBeforeTax, earningsAfterTax, ratio
}
