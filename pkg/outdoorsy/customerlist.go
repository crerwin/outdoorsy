package outdoorsy

import (
	"fmt"
	"sort"
	"strings"

	log "github.com/sirupsen/logrus"
)

type CustomerList struct {
	Customers []Customer
}

func (cl *CustomerList) addCustomer(customer Customer) {
	cl.Customers = append(cl.Customers, customer)
}

func (cl *CustomerList) Sort(sortBy string) {
	// TODO: unit test this
	fmt.Println(sortBy)
	if sortBy == "name" {
		sort.SliceStable(cl.Customers, func(i, j int) bool {
			// The requirements doc is a little unclear, so I'm assuming that
			// "sort by full name" means to sort by First+Last.  We do this by
			// concatenating the names and ignoring the space in between.
			return (strings.ToLower(
				cl.Customers[i].FirstName+cl.Customers[i].LastName) <
				strings.ToLower(
					cl.Customers[j].FirstName+cl.Customers[j].LastName))
		})
	} else if sortBy == "vehicle-type" {
		sort.SliceStable(cl.Customers, func(i, j int) bool {
			return (strings.ToLower(cl.Customers[i].Vehicle.VehicleType) <
				strings.ToLower(cl.Customers[j].Vehicle.VehicleType))
		})
	} else {
		log.Fatalf("Invalid --sort-by: %v", sortBy)
	}
}
