package kimsufi

import (
	"encoding/json"
	"github.com/WWWWaheb/goKimsufi/pkg/data"
	"github.com/sirupsen/logrus"
	"time"
)

var logger *logrus.Logger

var config kimsufiConfig

func StartBot(globalLogger *logrus.Logger, hwChan chan string, notifyChan chan string) {
	logger = globalLogger
	config = getConfig()
	go getKimsufiAvailability(hwChan, notifyChan)
}

func getKimsufiAvailability(hwChan chan string, notifyChan chan string) {

	for true {
		select {
		case hw := <-hwChan:
			logger.Println("Searching availability for server " + hw)
			config.Hardware = hw
			a := callApi()
			if a != "" {
				notifyChan <- a
			}
			break
		default:
			logger.Println("Searching availability for server " + config.Hardware)
			a := callApi()
			if a != "" {
				notifyChan <- a
			}
		}

		time.Sleep(config.PollInterval)
	}
}

func callApi() string {

	var a []data.Availabilities
	var str = ""
	body, err := QueryKimusfi(config)
	err = json.Unmarshal(body, &a)
	if err != nil {
		logger.Error("Error calling api : " + err.Error())
	}

	for _, v := range a {
		for _, vd := range v.Datacenters {
			if vd.Availability != "unavailable" {
				str = "Region: [" + v.Region + "]; Datacenter : [" + vd.Datacenter + "] Availability : [" + vd.Availability + "]"
			}
		}
	}
	logger.Info(str)
	return str
}
