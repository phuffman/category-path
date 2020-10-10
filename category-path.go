package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("Try: category-path in.csv out.txt")
		return
	}

	// Open the file (f)
	fin := os.Args[1]
	f, err := os.Open(fin)
	checkError("Cannot open file", err)
	defer f.Close()

	// Create empty map (m) to hold categories
	m := make(map[string]string)

	// Create reader (r) to read csv
	r := csv.NewReader(f)

	// Read csv to variable (rows)
	rows, err := r.ReadAll()
	checkError("Can't read file", err)

	// Iterate csv to add categories to map (m)
	for _, row := range rows {
		m[row[1]] = row[2]
	}

	// Create empty string slice (path) to hold all paths
	paths := []string{}

	// Iterate category csv (rows) to create a full path for each. Add all to string slice (paths).
	for _, row := range rows {

		// Create variables to hold current csv row
		categorytype := row[0]
		category := row[1]
		parent := row[2]

		// Create slice (path) to hold current path
		path := []string{}

		// 1) First, add category type
		path = append(path, categorytype)

		// 2) Then, add parent, parent of parent, etc.
		for parent != "" {
			path = append(path, parent)
			parent = m[parent]
		}
		// 3) Finally, add category
		path = append(path, category)

		// Add current 'path' to 'paths'
		paths = append(paths, strings.Join(path, " -> "))
	}

	// Sort paths
	sort.Strings(paths)

	// Convert slice to string (output)
	output := strings.Join(paths, "\n")

	// Print 'output' to screen
	println(output)

	// Create output file
	fout := os.Args[2]
	file, err := os.Create(fout)
	checkError("Can't create file", err)
	defer file.Close()

	// Write string to output file
	ln, err := io.WriteString(file, output)
	checkError("Can't write file", err)

	// ln (above) is required, and it needs to be used somewhere, so I'm showing it here ...
	fmt.Printf("String length: %v"+"\n", ln)
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
