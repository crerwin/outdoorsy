package outdoorsy

import "testing"

func TestNewPerson(t *testing.T) {
	testVehicle := newVehicle("Josephine", "sailboat", 25)
	testCustomer := newCustomer("Bob", "Sacamano", "bsacamano@kramerica.com", testVehicle)
	got := testCustomer.firstName
	want := "Bob"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = testCustomer.lastName
	want = "Sacamano"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

	got = testCustomer.email
	want = "bsacamano@kramerica.com"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
