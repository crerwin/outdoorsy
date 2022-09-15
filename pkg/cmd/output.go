package cmd

import (
	"fmt"

	"github.com/crerwin/outdoorsy/pkg/outdoorsy"
)

func outputCustomers(customers outdoorsy.CustomerList, sort bool, sortBy string) {
	if sort {
		customers.Sort(sortBy)
	}
	for _, customer := range customers.Customers {
		outputCustomer(customer)
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
