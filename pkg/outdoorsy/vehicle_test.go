package outdoorsy

import "testing"

func TestNewVehicle(t *testing.T) {
	testVehicle := newVehicle("Josephine", "sailboat", 25)
	got := testVehicle.Name
	want := "Josephine"
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = testVehicle.VehicleType
	want = "sailboat"
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	gotInt := testVehicle.Length
	wantInt := 25
	if gotInt != wantInt {
		t.Errorf("got %v, wanted %v", gotInt, wantInt)
	}
}
