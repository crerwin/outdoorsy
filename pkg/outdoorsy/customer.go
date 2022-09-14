package outdoorsy

type Customer struct {
	firstName string
	lastName  string
	email     string
	vehicle   *Vehicle
}

func newCustomer(firstName string, lastName string, email string, vehicle *Vehicle) *Customer {
	c := new(Customer)
	c.firstName = firstName
	c.lastName = lastName
	c.email = email
	c.vehicle = vehicle

	return c
}
