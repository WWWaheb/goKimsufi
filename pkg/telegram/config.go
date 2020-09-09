package telegram

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type telegramConfig struct {
	Token         string
	Hardware      []string
	ServerMapFile string
}

func getConfig() telegramConfig {
	token := os.Getenv("TELEGRAM_TOKEN")

	if token == "" {
		panic("No telegram token set up, please check your telegramConfig and set env TELEGRAM_TOKEN.")
	}

	s := os.Getenv("SERVER_MAP_FILE")

	if s == "" {
		panic("No server mapping json file provided. exiting.")
	}

	hwList, err := getHwList(s)

	if err != nil {
		panic("Please check the hardware mapping")
	}

	return telegramConfig{
		Token:         token,
		Hardware:      hwList,
		ServerMapFile: s,
	}
}

func getHwList(s string) ([]string, error) {
	data, err := ioutil.ReadFile(s)
	if err != nil {
		return nil, err
	}
	var hardwareMap map[string]string
	err = json.Unmarshal(data, &hardwareMap)
	if err != nil {
		return nil, err
	}
	hwList := make([]string, 0)
	for _, v := range hardwareMap {
		hwList = append(hwList, v)
	}

	return hwList, nil
}
