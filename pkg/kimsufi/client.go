package kimsufi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func QueryKimusfi(c kimsufiConfig) ([]byte, error) {
	var a []byte

	kimsufiUrl, err := url.Parse(c.KimsufiUrl)
	if err != nil {
		return a, err
	}

	hwCode, err := getHardwareCode(c.Hardware)
	if err != nil {
		return a, err
	}
	q := kimsufiUrl.Query()
	q.Set("country", c.Country)
	q.Set("hardware", hwCode)
	kimsufiUrl.RawQuery = q.Encode()

	req, err := http.Get(kimsufiUrl.String())
	if err != nil {
		return a, err
	}

	body, err := ioutil.ReadAll(req.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func getHardwareCode(h string) (string, error) {

	data, err := ioutil.ReadFile(config.ServerMapFile)
	if err != nil {
		logger.Error(err)
		return "", err
	}
	var hardwareMap map[string]interface{}
	err = json.Unmarshal(data, &hardwareMap)
	if err != nil {
		return "", err
	}
	for k, v := range hardwareMap {
		if h == v {
			return k, nil
		}
	}
	panic("Hardware type " + h + " not found.")
}
