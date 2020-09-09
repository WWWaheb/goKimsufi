package kimsufi

import (
	"os"
	"strconv"
	"time"
)

type kimsufiConfig struct {
	KimsufiUrl    string
	Country       string
	Hardware      string
	PollInterval  time.Duration
	ServerMapFile string
}

func getConfig() kimsufiConfig {

	url := os.Getenv("KIMSUFI_URL")
	c := os.Getenv("KIMSUFI_COUNTRY")
	hardware := os.Getenv("KIMSUFI_HARDWARE")
	pollValue, err := strconv.Atoi(os.Getenv("KIMSUFI_POLLTIME"))
	s := os.Getenv("SERVER_MAP_FILE")

	if err != nil {
		pollValue = 10
	}

	if url == "" {
		url = "https://ca.ovh.com/engine/api/dedicated/server/availabilities?"
	}

	if c == "" {
		c = "FR"
	}

	if hardware == "" {
		hardware = "KS-11"
	}

	if s == "" {
		panic("No server mapping json file provided. exiting.")
	}

	config := kimsufiConfig{
		KimsufiUrl:    url,
		Country:       c,
		Hardware:      hardware,
		PollInterval:  time.Duration(pollValue) * time.Second,
		ServerMapFile: s,
	}

	return config
}
