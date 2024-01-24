package prices

import (
	"fmt"

	"github.com/gillgillint/price-calculator/conversion"
	"github.com/gillgillint/price-calculator/iomanager"
)

type PricesIncludedTaxJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	PricesIncludedTax map[string]string   `json:"prices_included_tax"`
}

func (job *PricesIncludedTaxJob) LoadData() error {
	lines, err := job.IOManager.ReadLines()

	if err != nil {
		fmt.Println(err)
		return err
	}

	prices, err := conversion.StringsToFloats(lines)

	if err != nil {
		fmt.Println(err)
		return err
	}

	job.InputPrices = prices
	return nil
}

func (job *PricesIncludedTaxJob) Process(doneChan chan bool, errorChan chan error) {
	err := job.LoadData()

	if err != nil {
		// return err
		errorChan <- err
		return
	}
	result := make(map[string]string)

	for _, price := range job.InputPrices {
		PricesIncludedTax := price * (1 + job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", PricesIncludedTax)
	}

	job.PricesIncludedTax = result
	job.IOManager.WriteResult(job)
	doneChan <- true
}

func NewPricesIncludedTaxJob(iom iomanager.IOManager, taxRate float64) *PricesIncludedTaxJob {
	return &PricesIncludedTaxJob{
		IOManager:   iom,
		TaxRate:     taxRate,
		InputPrices: []float64{10, 20, 30},
	}
}
