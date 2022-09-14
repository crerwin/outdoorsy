package outdoorsy

type Vehicle struct {
	name        string
	vehicleType string
	length      int
}

func newVehicle(name string, vehicleType string, length int) *Vehicle {
	v := new(Vehicle)
	v.name = name
	v.vehicleType = vehicleType
	v.length = length

	return v
}
