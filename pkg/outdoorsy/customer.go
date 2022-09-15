package outdoorsy

type Customer struct {
	FirstName string
	LastName  string
	Email     string
	Vehicle   *Vehicle
}

func newCustomer(firstName string, lastName string, email string, vehicle *Vehicle) *Customer {
	c := new(Customer)
	c.FirstName = firstName
	c.LastName = lastName
	c.Email = email
	c.Vehicle = vehicle

	return c
}
