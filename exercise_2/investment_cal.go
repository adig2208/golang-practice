package main

import (
	"fmt"
	"math"
)

func main() {
	const inflation = 2.5
	var investmentAmount float64
	var years float64
	expectedReturnRate := 5.5

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	var futureValue = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	var futureRealValue = futureValue / math.Pow(1+inflation/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)
}

func getValues(investmentAmount float64, expectedReturnRate float64, years float64) {

}
