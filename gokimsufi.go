package main

import (
	client "github.com/WWWWaheb/goKimsufi/pkg/client"
	conf "github.com/WWWWaheb/goKimsufi/pkg/config"
)

func main() {
	config := conf.GetConfig()
	client.QueryKimusfi(config)
}
