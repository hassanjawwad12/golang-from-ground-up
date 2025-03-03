package filemanagement

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloattoFile(value float64, filename string) {
	valueText := fmt.Sprintf("%f", value)
	os.WriteFile(filename, []byte(valueText), 0644)
}

func ReadFloatFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading balance file", err)
		return 1000, errors.New("error reading balance file")
	}
	value, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		fmt.Println("Error parsing balance", err)
		return 1000, errors.New("error parsing balance")
	}
	return value, nil
}
