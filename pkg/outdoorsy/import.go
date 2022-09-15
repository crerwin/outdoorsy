package outdoorsy

import (
	"bufio"
	"errors"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

func readFile(filename string) []byte {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file %v with error %v", filename, err)
	}

	return content
}

func loadFile(filename string) []Customer {
	var customers []Customer
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error opening file %v: %v", filename, err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		rowItems, err := splitRow(scanner.Text())
		if err != nil {
			log.Fatalf("Error splitting row %v: %v", scanner.Text(), err)
		}
		customers = append(customers, loadRow(rowItems))
	}
	return customers
}

func LoadFiles(filenames []string) []Customer {
	var customers []Customer
	for _, filename := range filenames {
		customers = append(customers, loadFile(filename)...)
	}

	return customers
}

func splitRow(row string) ([]string, error) {
	var delimiter string

	if strings.Contains(row, ",") && strings.Contains(row, "|") {
		errorMsg := "Bad row.  Both , and | found in row: " + row
		return nil, errors.New(errorMsg)
	} else if strings.Contains(row, ",") {
		log.Debug("Comma delimited row detected")
		delimiter = ","
	} else if strings.Contains(row, "|") {
		log.Debug("Pipe delimited row detected")
		delimiter = "|"
	} else {
		errorMsg := "No valid delimiter found in row: " + row
		return nil, errors.New(errorMsg)
	}
	items := strings.Split(row, delimiter)
	return items, nil
}

func loadRow(items []string) Customer {
	firstName := items[0]
	lastName := items[1]
	email := items[2]
	vehicleType := items[3]
	vehicleName := items[4]
	vehicleLength := 5
	return *newCustomer(
		firstName, lastName, email, newVehicle(
			vehicleName, vehicleType, vehicleLength,
		),
	)
}
