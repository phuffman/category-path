package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Please provide a filename")
		return
	}

	// Open the file (f)
	filename := os.Args[1]
	f, err := os.Open(filename)
	checkError("Cannot open file", err)
	defer f.Close()

	// Create new map (m)
	m := make(map[string]string)

	// Start the reader (r)
	r := csv.NewReader(f)

	// Read the csv into variable
	rows, err := r.ReadAll()
	checkError("Cannot read file", err)

	// Populate the map
	for _, row := range rows {
		m[row[1]] = row[2]
	}

	lines := []string{}

	// Lookup
	for _, row := range rows {

		categorytype := row[0]
		category := row[1]
		parent := row[2]

		path := []string{categorytype}

		for parent != "" {
			path = append(path, parent)
			parent = m[parent]
		}

		path = append(path, category)
		lines = append(lines, strings.Join(path, " -> "))
	}
	sort.Strings(lines)

	for _, line := range lines {
		println(line)
	}

	file, err := os.Create("result.csv")
	checkError("Cannot create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, value := range data {
		err := writer.Write(value)
		checkError("Cannot write to file", err)
	}

}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
