package cmdmanager

import "fmt"

type CMDmanager struct {
}

func (cmd CMDmanager) ReadLines() ([]string, error) {
	fmt.Println("Please enter Prices. Confirm each with enter.")
	var prices []string
	for {
		var price string
		fmt.Print("Price")
		fmt.Scan(&price)
		if price == "" {
			break
		}
		prices = append(prices, price)
	}
	return prices, nil
}

func (cmd CMDmanager) WriteJSON(data interface{}) error {
	fmt.Println(data)
	return nil
}

func New() CMDmanager {
	return CMDmanager{}
}
