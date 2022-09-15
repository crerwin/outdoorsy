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

func ReadFiles(filenames []string) string {
	var content string
	for _, filename := range filenames {
		content += string(readFile(filename))
		content += "\n"
	}

	scanner := bufio.NewScanner(strings.NewReader(content))
	for scanner.Scan() {
		items, err := splitRow(scanner.Text())
		if err != nil {
			log.Fatalf("Error splitting row %v: %v", scanner.Text(), err)
		} else {
			loadRow(items)
		}

	}
	return content
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
