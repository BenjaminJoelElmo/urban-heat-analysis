package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

// CO2Record is defined ONCE in the main package
type CO2Record struct {
	CountryName  string
	CountryCode  string
	Year         int
	CO2Emissions float64
}

func LoadAndClean(filePath string) ([]CO2Record, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	_, _ = reader.Read() // Skip header

	var records []CO2Record
	for i := 2; ; i++ {
		row, err := reader.Read()
		if err != nil {
			break
		}

		year, _ := strconv.Atoi(row[4])
		co2, _ := strconv.ParseFloat(row[5], 64)

		records = append(records, CO2Record{
			CountryName:  row[0],
			CountryCode:  row[1],
			Year:         year,
			CO2Emissions: co2,
		})
	}
	return records, nil
}
