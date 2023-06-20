package main

import (
	"encoding/csv"
	"fmt"
	"strings"
)

func main() {
	// Sample data for file1.csv
	file1Data := `John,Doe,123 Main St
Jane,Smith,456 Oak Ave
Bob,Johnson,789 Elm St`

	// Sample data for file2.csv
	file2Data := `Jane,Smith,456 Oak Ave
Bob,Johnson,789 Elm St
Sally,Jones,321 Pine St`

	// Read the contents of the first CSV file
	reader1 := csv.NewReader(strings.NewReader(file1Data))
	records1, err := reader1.ReadAll()
	if err != nil {
		panic(err)
	}

	// Read the contents of the second CSV file
	reader2 := csv.NewReader(strings.NewReader(file2Data))
	records2, err := reader2.ReadAll()
	if err != nil {
		panic(err)
	}

	// Compare the two CSV files for matches
	for _, record1 := range records1 {
		for _, record2 := range records2 {
			if record1[0] == record2[0] && record1[1] == record2[1] {
				fmt.Printf("Match found: %v\n", record1)
			}
		}
	}
}
