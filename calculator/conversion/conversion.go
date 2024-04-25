package conversion

import (
	"errors"
	"fmt"
	"strconv"
)

func StringToFLoat(inString []string) ([]float64, error) {
	var outValues []float64
	for _, val := range inString {
		floatPrice, err := strconv.ParseFloat(val, 64)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("Error!")
		}
		outValues = append(outValues, floatPrice)
	}
	return outValues, nil
}
