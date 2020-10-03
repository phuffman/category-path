package main

import (
	"encoding/csv"
	"os"
	"sort"
	"strings"

	"github.com/ogier/pflag"
)

// flags
var (
	fin string
)

func main() {
	pflag.Parse()

	// Create new map (m)
	m := make(map[string]string)

	// Open the file (f)
	f, err := os.Open(os.Args[1])
	if err != nil {
		panic(err)
	}

	// Start the reader (r)
	r := csv.NewReader(f)

	// Read the csv into variable
	rows, err := r.ReadAll()
	if err != nil {
		panic(err)
	}

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

}

func init() {
	pflag.StringVarP(&fin, "inputfile", "i", "", "Input File")
}
