package configuration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Configuration represents hte elements in the config
type Configuration struct {
	APIKey string
}

// Config returns a Configuration with values read from conf.json
func Config(fileName string) (*Configuration, error) {
	b, err := ioutil.ReadFile(fileName)
	l := &Configuration{}

	if err != nil {
		return l, err
	}
	err = json.Unmarshal(b, &l)
	if err != nil {
		fmt.Print("bad json ", err)
	}

	return l, err
}
