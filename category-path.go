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

	// Verify input has 2 args
	if len(os.Args) != 2 {
		fmt.Println("Try: category-path in.csv")
		return
	}

	// Open file (f)
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal("Can't open file", err)
	}
	defer f.Close()

	// Create a reader (r)
	r := csv.NewReader(f)

	// Get header row (header)
	header, err := r.Read()
	if err != nil {
		log.Fatal("Can't open header", err)
	}

	// Get column numbers of key fields

	// Create struct (Columns)
	type Columns struct {
		categorytype int
		parent       int
		category     int
	}

	// Create Columns variable (cols)
	cols := Columns{
		categorytype: -1,
		parent:       -1,
		category:     -1,
	}

	// Get column number of name, parent and type columns (-1 if not found)
	i := 0
	for _, column := range header {
		switch {
		case strings.Contains(column, "Type"):
			cols.categorytype = i
		case strings.Contains(column, "Parent"):
			cols.parent = i
		case strings.Contains(column, "Name"):
			cols.category = i
		}
		i++
	}

	// Verify category and parent cateory exist
	if cols.category == -1 || cols.parent == -1 {
		fmt.Println("You need columns called 'Category Name' and 'Parent Category'.")
		return
	}

	// Create empty map (m) to hold categories
	m := make(map[string]string)

	// Get remaining rows (rows)
	rows, err := r.ReadAll()
	if err != nil {
		log.Fatal("Can't read file", err)
	}

	// Iterate csv, adding categories to map (m)
	for _, row := range rows {
		m[row[cols.category]] = row[cols.parent]
	}

	// Create empty string slice (path) to hold all paths
	paths := []string{}

	// Iterate csv again to create a full path for each.
	// Add each path to slice (paths).
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

		// 3) Finally, prepend category type (if there is one)
		if cols.categorytype != -1 {
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

	// 'output' to screen
	// println(output)

	// Create output file (out.csv)
	file, err := os.Create("out.csv")
	if err != nil {
		log.Fatal("Can't create file", err)
	}
	defer file.Close()

	// Write to output file
	_, err = io.WriteString(file, output)
	if err != nil {
		log.Fatal("Can't write file", err)
	}
}
