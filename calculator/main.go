package main

import (
	"fmt"

	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
)

func main() {

	taxRates := []float64{0.0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	for i, rate := range taxRates {
		doneChans[i] = make(chan bool)
		fm := filemanager.New("prices.txt", fmt.Sprintf("results_%.0f.json", rate*100))
		// cmdm := cmdmanager.New()
		priceJob := prices.New(fm, rate)
		go priceJob.Process(doneChans[i])

	}
	for _, doneChan := range doneChans {
		<-doneChan
	}
}
