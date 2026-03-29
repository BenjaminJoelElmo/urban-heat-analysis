package main

import (
	"encoding/csv"
	"os"
	"strconv"
)

// SaveCleaned writes processed data to CSV
func SaveCleaned(data []CO2Record, outputPath string) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	writer.Write([]string{"CountryName", "CountryCode", "Year", "CO2Emissions"})

	// Write data
	for _, d := range data {
		writer.Write([]string{
			d.CountryName,
			d.CountryCode,
			strconv.Itoa(d.Year),
			strconv.FormatFloat(d.CO2Emissions, 'f', 2, 64),
		})
	}
	return nil
}
