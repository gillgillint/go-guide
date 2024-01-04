package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	// fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	// fmt.Print("years: ")
	outputText("years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculatorFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	// futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)
	formattedRFV := fmt.Sprintf("FutureValue (adjusted for Inflation): %.1f\n", futureRealValue)

	// outputs
	// fmt.Println("futureValue: ", futureValue)
	// fmt.Println(" futureValue (adjusted for Inflation):", futureRealValue)
	// fmt.Printf(`Future Value: %.1f
	// Future Value (adjusted for inflation): %.1f`, futureValue, futureRealValue)
	fmt.Print(formattedFV, formattedRFV)

}

func outputText(text string) {
	fmt.Print(text)
}

func calculatorFutureValues(investmentAmount float64, expectedReturnRate float64, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
	// return

}
