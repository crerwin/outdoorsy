package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/crerwin/outdoorsy/pkg/outdoorsy"
	"github.com/olekukonko/tablewriter"
	log "github.com/sirupsen/logrus"
)

func outputCustomers(customers outdoorsy.CustomerList, sort bool, sortBy string) {
	if sort {
		log.Debug("Sort flag set - sorting customers")
		customers.Sort(sortBy)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{
		"Name", "Email", "Vehicle Type", "Vehicle Name", "Vehicle Length",
	})
	for _, customer := range customers.Customers {
		table.Append(customerToStrings(customer))
	}
	table.Render()

	fmt.Printf("Total customer records: %v\n", len(customers.Customers))
}

// maybe this should live inside the outdoorsy package?
func customerToStrings(customer outdoorsy.Customer) []string {
	return []string{
		customer.FirstName + " " + customer.LastName,
		customer.Email,
		customer.Vehicle.VehicleType,
		customer.Vehicle.Name,
		strconv.Itoa(customer.Vehicle.Length),
	}
}
