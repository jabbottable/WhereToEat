package configuration

import "testing"

func TestConfig(t *testing.T) {
	for _, c := range []struct {
		in, want string
		hasError bool
	}{
		{"testdata/confTest.json", "TestApiKey", false},
		{"testdata/doesNotExist.json", "", true},
		{"testdata/confTestNotJson.json", "", true},
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
