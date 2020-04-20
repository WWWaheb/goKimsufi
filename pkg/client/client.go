package client

import (
	"encoding/json"
	"fmt"
	c "github.com/WWWWaheb/goKimsufi/pkg/config"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Datacenter struct {
	DataCenter   string `json:"datacenter"`
	Availability string `json:"availability"`
}

type Availabilities struct {
	Region      string       `json:"region"`
	Hardware    string       `json:"hardware"`
	Datacenters []Datacenter `json:"datacenter"`
}

func QueryKimusfi(c c.Config) (Availabilities, error) {
	kimsufiUrl, err := url.Parse(c.KimsufiUrl)
	if err != nil {
		return Availabilities{}, err
	}
	q := kimsufiUrl.Query()
	q.Set("country", c.Country)
	q.Set("instanceType", c.InstanceType)
	kimsufiUrl.RawQuery = q.Encode()

	req, err := http.Get(kimsufiUrl.String())
	if err != nil {
		return Availabilities{}, err
	}
	var a []Availabilities

	body, err := ioutil.ReadAll(req.Body)

	err = json.Unmarshal(body, &a)
	if err != nil {
		return Availabilities{}, err
	}

	fmt.Println(a[0])

	return Availabilities{}, nil
}
