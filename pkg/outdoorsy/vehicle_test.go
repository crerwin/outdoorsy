package outdoorsy

import "testing"

func TestNewVehicle(t *testing.T) {
	testVehicle := newVehicle("Josephine", "sailboat", 25)
	got := testVehicle.name
	want := "Josephine"
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = testVehicle.vehicleType
	want = "sailboat"
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	gotInt := testVehicle.length
	wantInt := 25
	if gotInt != wantInt {
		t.Errorf("got %v, wanted %v", gotInt, wantInt)
	}
}
