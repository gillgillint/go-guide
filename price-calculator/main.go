package main

import (
	"fmt"

	"github.com/gillgillint/price-calculator/cmdmanager"
	"github.com/gillgillint/price-calculator/prices"
)

func main() {
	taxRates := []float64{0, 0.07, 0.1, 0.15}
	dones := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		dones[index] = make(chan bool)
		errorChans[index] = make(chan error)
		// fm := filemanager.New("prices.txt", fmt.Sprintf("result_%.0f.json", taxRate*100))
		cmdm := cmdmanager.New()
		priceJob := prices.NewPricesIncludedTaxJob(cmdm, taxRate)
		go priceJob.Process(dones[index], errorChans[index])

		// if err != nil {
		// 	fmt.Println("Could not pass job")
		// 	fmt.Println(err)
		// }
	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Println(err)
			}
		case <-dones[index]:
			fmt.Println("Done!")
		}
	}

}
