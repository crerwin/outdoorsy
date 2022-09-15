package cmd

import (
	"fmt"

	"github.com/crerwin/outdoorsy/pkg/outdoorsy"
)

func outputCustomers(customers []outdoorsy.Customer) {
	for _, customer := range customers {
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
