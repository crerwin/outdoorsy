package outdoorsy

import (
	"reflect"
	"testing"
)

func TestSplitRow(t *testing.T) {
	cases := []struct {
		row  string
		want []string
	}{
		{"a,b", []string{"a", "b"}},
		{"a|b", []string{"a", "b"}},
	}
	for _, c := range cases {
		got, _ := splitRow(c.row)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got %v from splitRow(%v), expected %v", got, c.row, c.want)
		}
	}
}

func TestSplitRowBadData(t *testing.T) {
	cases := []struct {
		row string
		// expected error message
		want string
	}{
		{"ab", "No valid delimiter found in row: ab"},
		{"a,b|c", "Bad row.  Both , and | found in row: a,b|c"},
	}

	for _, c := range cases {
		got, err := splitRow(c.row)
		if got != nil {
			t.Errorf("Unexpectedly got return value %v from splitRow(%v) when we expected the error %v", err, c.row, c.want)
		} else if err.Error() != c.want {
			t.Errorf("Expected error message %v from splitRow(%v) but received %v", c.want, c.row, err)
		}
	}
}
