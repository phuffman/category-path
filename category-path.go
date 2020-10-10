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
		fmt.Println("Use the category-path file-in file-out")
		return
	}

	// Open the file (f)
	filename := os.Args[1]
	f, err := os.Open(filename)
	checkError("Cannot open file", err)
	defer f.Close()

	// Create map to hold categories (m)
	m := make(map[string]string)

	// Create reader (r)
	r := csv.NewReader(f)

	// Read csv to variable (rows)
	rows, err := r.ReadAll()
	checkError("Can't read file", err)

	// Add categories to map (m)
	for _, row := range rows {
		m[row[1]] = row[2]
	}

	// Create slice (path) to hold all paths
	paths := []string{}

	// Iterate csv (rows)
	for _, row := range rows {

		// Create variables to hold current csv row
		categorytype := row[0]
		category := row[1]
		parent := row[2]

		// Create slice (path) to hold current path
		path := []string{}

		// 1) Add category type to 'path'
		path = append(path, categorytype)

		// 2) Add parent to 'path', then loop adding parent of parent, etc. to 'path'
		for parent != "" {
			path = append(path, parent)
			parent = m[parent]
		}
		// 3) Finally, add category to 'path'
		path = append(path, category)

		// Add 'path' to 'paths'
		paths = append(paths, strings.Join(path, " -> "))
	}

	// Sort paths
	sort.Strings(paths)

	output := strings.Join(paths, "\n")

	println(output)

	/*
		for _, line := range paths {
			println(line)
		}
	*/

	file, err := os.Create("./out.csv")
	checkError("Can't create file", err)
	defer file.Close()

	ln, err := io.WriteString(file, output)
	checkError("Can't write file", err)

	fmt.Printf("String length: %v"+"\n", ln)
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
