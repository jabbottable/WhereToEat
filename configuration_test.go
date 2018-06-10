package wheretoeat

import "testing"

func TestConfig(t *testing.T) {
	for _, c := range []struct {
		in, want string
		hasError bool
	}{
		{"confTest.json", "TestApiKey", false},
		{"doesNotExist.json", "", true},
		{"confTestNotJson.json", "", true},
	} {
		got, err := Config(c.in)
		if got.APIKey != c.want {
			t.Errorf("Config(%q) == %q, want %q", c.in, got.APIKey, c.want)
		}

		if (err != nil) != c.hasError {
			t.Errorf("Config(%q) has error %t, hasError %t", c.in, err != nil, c.hasError)
		}
	}
}
