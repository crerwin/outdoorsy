package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/crerwin/outdoorsy/pkg/outdoorsy"
	"github.com/olekukonko/tablewriter"
)

func outputCustomers(customers outdoorsy.CustomerList, sort bool, sortBy string) {
	if sort {
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

func outputCustomer(customer outdoorsy.Customer) {
	fmt.Println(
		customer.FirstName,
		customer.LastName,
		customer.Email,
		customer.Vehicle.VehicleType,
		customer.Vehicle.Name,
		customer.Vehicle.Length,
	)
}
