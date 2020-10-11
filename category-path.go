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

	// Check command line for 'category-path in.csv out.txt'
	if len(os.Args) != 2 {
		fmt.Println("Try: category-path in.csv")
		return
	}

	// Open the file (f)
	f, err := os.Open(os.Args[1])
	checkError("Cannot open file", err)
	defer f.Close()

	// Create a reader (r)
	r := csv.NewReader(f)

	// Read the first row (header)
	header, err := r.Read()
	checkError("Can't open header", err)

	type Columns struct {
		categorytype, parent, category int
	}

	cols := Columns{0, 0, 100}

	i := 0
	for _, column := range header {
		switch {
		case strings.Contains(column, "Category Type"):
			cols.categorytype = i
		case strings.Contains(column, "Parent Category"):
			cols.parent = i
		case strings.Contains(column, "Category Name"):
			cols.category = i
		}
		i++
	}

	// Verify Category and Parent Cateory columns exist
	if cols.category == 0 || cols.parent == 0 {
		fmt.Println("I need columns called 'Category Name' and 'Parent Category'.")
		return
	}

	// Create empty map (m) to hold categories
	m := make(map[string]string)

	// Read remaining rows to variable (rows)
	rows, err := r.ReadAll()
	checkError("Can't read file", err)

	// Iterate csv to add categories to map (m)
	for _, row := range rows {
		m[row[cols.category]] = row[cols.parent]
	}

	// Create empty string slice (path) to hold all paths
	paths := []string{}

	// Iterate category csv (rows) to create a full path for each. Add all to string slice (paths).
	for _, row := range rows {

		// Create variables to hold parent and category from current row
		parent := row[cols.parent]
		category := row[cols.category]

		// Create slice (path) to hold current path
		path := []string{}

		// 1) First, add category
		path = append(path, category)

		// 2) Then, prepend parent, parent of parent, etc.
		for parent != "" {
			path = append([]string{parent}, path...)
			parent = m[parent]
		}

		// 3) Finally, if exists, prepend category type
		if cols.categorytype != 100 {
			categorytype := row[cols.categorytype]
			path = append([]string{categorytype}, path...)
		}

		// Add current 'path' to 'paths'
		paths = append(paths, strings.Join(path, " > "))
	}

	// Sort paths
	sort.Strings(paths)

	// Convert slice to string (output)
	output := strings.Join(paths, "\n")

	// Print 'output' to screen
	println(output)

	// Create output file
	file, err := os.Create("out.csv")
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
