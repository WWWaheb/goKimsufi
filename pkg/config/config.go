package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	KimsufiUrl   string
	Country      string
	InstanceType string
}

type Instance struct {
	instanceType string
}

func GetConfig() Config {

	url := os.Getenv("KIMSUFI_URL")
	c := os.Getenv("KIMSUFI_COUNTRY")
	instance := os.Getenv("KIMSUFI_INSTANCE_TYPE")

	if url == "" {
		url = "https://ca.ovh.com/engine/api/dedicated/server/availabilities?"
	}

	if c == "" {
		c = "fra"
	}

	if instance == "" {
		instance = "KS-1"
	}

	config := Config{
		KimsufiUrl:   url,
		Country:      c,
		InstanceType: instance,
	}

	return config
}

func getInstanceName(instance string) string {

	data, err := ioutil.ReadFile("./conf/server_mapping.json")
	if err != nil {
		fmt.Print(err)
	}
	instanceMapping := []Instance{}
	json.Unmarshal([]byte(data), &instanceMapping)
	fmt.Println(instanceMapping[0])
	return "1804sk12"
}
