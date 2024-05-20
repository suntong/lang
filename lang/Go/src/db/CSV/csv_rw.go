package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Person represents a single record in the CSV file
type Person struct {
	Name string
	Age  int
	City string
}

// readCSV reads the data from a CSV file into a slice of Person structs
func readCSV(filename string) ([]Person, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("could not open CSV file: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("could not read CSV file: %v", err)
	}

	var people []Person
	for i, record := range records {
		// Skip the header
		if i == 0 {
			continue
		}
		age, err := strconv.Atoi(record[1])
		if err != nil {
			return nil, fmt.Errorf("could not convert age to int: %v", err)
		}
		person := Person{
			Name: record[0],
			Age:  age,
			City: record[2],
		}
		people = append(people, person)
	}
	return people, nil
}

// writeCSV writes the data from a slice of Person structs to a CSV file
func writeCSV(filename string, people []Person) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("could not create CSV file: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	err = writer.Write([]string{"Name", "Age", "City"})
	if err != nil {
		return fmt.Errorf("could not write header to CSV file: %v", err)
	}

	for _, person := range people {
		record := []string{person.Name, strconv.Itoa(person.Age), person.City}
		err = writer.Write(record)
		if err != nil {
			return fmt.Errorf("could not write record to CSV file: %v", err)
		}
	}
	return nil
}

func main() {
	// Read CSV file
	people, err := readCSV("csv_rw.input.csv")
	if err != nil {
		log.Fatalf("Error reading CSV file: %v", err)
	}

	// Process data (e.g., increment age by 1)
	for i := range people {
		people[i].Age++
	}

	// Write processed data to a new CSV file
	err = writeCSV("csv_rw.output.csv", people)
	if err != nil {
		log.Fatalf("Error writing CSV file: %v", err)
	}

	fmt.Println("CSV data successfully read, processed, and written to output.csv")
}
