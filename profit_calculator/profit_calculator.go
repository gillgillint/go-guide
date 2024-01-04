package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// var revenue float64
	// var expenses float64
	// var taxRate float64

	// fmt.Print("Revenue: ")
	// fmt.Scan(&revenue)
	revenue, err := getUserInput("Revenue: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Print("Expenses: ")
	// fmt.Scan(&expenses)
	expenses, err := getUserInput("Expenses: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	// fmt.Print("Tax Rate: ")
	// fmt.Scan(&taxRate)
	taxRate, err := getUserInput("Tax Rate: ")

	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateFinancial(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)

	storeResults(ebt, profit, ratio)
}

func calculateFinancial(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value must be a positive number.")
	}

	return userInput, nil
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT: %.1f\nProfit: %.1f\nRatio: %.3f\n", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(results), 0644)
}
