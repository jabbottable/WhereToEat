package helper
import (
    "encoding/json"
    "fmt"
    "io/ioutil"
)

type Configuration struct {
    ApiKey    string
}

func Config() (*Configuration, error) {
        b, err := ioutil.ReadFile("conf.json")
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
