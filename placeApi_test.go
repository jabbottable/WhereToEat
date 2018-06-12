package wheretoeat

import "testing"

func TestLocationString(t *testing.T) {
	for _, c := range []struct {
		in   Location
		want string
	}{
		{
			Location{
				Latitude:  2121.0,
				Longitude: 12121.0,
			},
			"2121.0000000,12121.0000000",
		},
		{
			Location{
				Latitude:  2121.1234567,
				Longitude: 12121.7654321,
			},
			"2121.1234567,12121.7654321",
		},
		{
			Location{
				Latitude:  -2121.1234567,
				Longitude: -12121.7654321,
			},
			"-2121.1234567,-12121.7654321",
		},
		{
			Location{
				Latitude:  2121.12345678910111213141516,
				Longitude: 12121.16151413121110987654321,
			},
			"2121.1234568,12121.1615141",
		},
		{
			Location{},
			"0.0000000,0.0000000",
		},
	} {
		got := c.in.LocationString()
		if got != c.want {
			t.Errorf("LocationString(%f,%f) == %q, want %q", c.in.Latitude, c.in.Longitude, got, c.want)
		}
	}
}
