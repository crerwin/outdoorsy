package outdoorsy

type Vehicle struct {
	Name        string
	VehicleType string
	Length      int
}

func newVehicle(name string, vehicleType string, length int) *Vehicle {
	v := new(Vehicle)
	v.Name = name
	v.VehicleType = vehicleType
	v.Length = length

	return v
}
