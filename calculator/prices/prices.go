package prices

import (
	"fmt"

	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
)

type TaxIncludedPriceJob struct {
	IOmanager        iomanager.IOManager `json: "-"`
	TaxRate          float64             `json: "tax_rate"`
	InputPrices      []float64           `json: "input_prices"`
	TaxIncludedPrice map[string]string   `json: "tax_included_prices"`
}

func (job *TaxIncludedPriceJob) LoadData() {
	lines, err := job.IOmanager.ReadLines()
	if err != nil {
		fmt.Println(err)
		return
	}
	prices, err := conversion.StringToFLoat(lines)
	if err != nil {
		fmt.Println(err)
		return
	}

	job.InputPrices = prices
}

func (job *TaxIncludedPriceJob) Process(doneChan chan bool) {
	job.LoadData()
	results := make(map[string]string)
	for _, price := range job.InputPrices {
		taxIncludedString := price * (1 + job.TaxRate)
		results[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxIncludedString)
	}
	job.TaxIncludedPrice = results
	job.IOmanager.WriteJSON(job)
	doneChan <- true

}
func New(iom iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOmanager:   iom,
		InputPrices: []float64{10, 20, 30},
		TaxRate:     taxRate,
	}

}
