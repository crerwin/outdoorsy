package outdoorsy

import (
	"bufio"
	"errors"
	"os"
	"strconv"
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

func loadFile(filename string) {
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
		OutdoorsyCustomers.addCustomer(loadRow(rowItems))
	}
}

func LoadFiles(filenames []string) []Customer {
	var customers []Customer
	for _, filename := range filenames {
		loadFile(filename)
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
	vehicleLength := processLength(items[5])
	return *newCustomer(
		firstName, lastName, email, newVehicle(
			vehicleName, vehicleType, vehicleLength,
		),
	)
}

func processLength(lengthString string) int {
	var length int
	var err error

	if _, err2 := strconv.Atoi(lengthString); err2 == nil {
		// 32
		length, err = strconv.Atoi(lengthString)
	} else if strings.HasSuffix(lengthString, "'") {
		// 32'
		length, err = strconv.Atoi(lengthString[:len(lengthString)-1])
	} else if strings.HasSuffix(lengthString, "’") {
		// 32’
		log.Debug("Dealing with unicode character in vehicle length")

		runes := []rune(lengthString)
		length, err = strconv.Atoi(string(runes[:len(runes)-1]))
	} else if strings.HasSuffix(lengthString, " feet") {
		// 32 feet
		length, err = strconv.Atoi(lengthString[:len(lengthString)-5])
	} else if strings.HasSuffix(lengthString, " ft") {
		// 32 ft
		length, err = strconv.Atoi(lengthString[:len(lengthString)-3])
	} else {
		log.Fatalf("Unexpected vehicle length format: %v", lengthString)
	}

	if err != nil {
		log.Fatalf("Error getting integer from %v: %v", lengthString, err)
	}

	return length
}
