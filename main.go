package main

import (
	"log"
	"os"
	"path/filepath"
)

func main() {
	cleanedData, err := LoadAndClean("data/raw/NGDS_CO2_Germany.csv")
	if err != nil {
		log.Fatalf("❌ Failed to load data: %v", err)
	}

	log.Printf("✅ Successfully loaded %d records", len(cleanedData))
	log.Println("First 3 records:")
	for i, r := range cleanedData[:3] {
		log.Printf("[%d] %s (%s) | %d | %.1f Mt CO2e",
			i+1, r.CountryName, r.CountryCode, r.Year, r.CO2Emissions)
	}

	// Run trend analysis
	AnalyzeTrend(cleanedData)

	// Create directory if missing
	outputPath := "data/cleaned/germany_co2_cleaned.csv"
	dir := filepath.Dir(outputPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("❌ Failed to create directory: %v", err)
	}

	// Save processed data
	if err := SaveCleaned(cleanedData, outputPath); err != nil {
		log.Fatalf("❌ Failed to save cleaned data: %v", err)
	}
}
