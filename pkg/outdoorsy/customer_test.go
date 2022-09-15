package outdoorsy

import "testing"

func TestNewPerson(t *testing.T) {
	testVehicle := newVehicle("Josephine", "sailboat", 25)
	testCustomer := newCustomer("Bob", "Sacamano", "bsacamano@kramerica.com", testVehicle)
	got := testCustomer.FirstName
	want := "Bob"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = testCustomer.LastName
	want = "Sacamano"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = testCustomer.Email
	want = "bsacamano@kramerica.com"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
