package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
	"github.com/junzh0u/smartthings"
)

func main() {
	flag.Set("stderrthreshold", "FATAL")
	domain := flag.String("domain", "graph-na02-useast1.api.smartthings.com", "SmartThings domain")
	username := flag.String("username", "", "SmartThings username")
	password := flag.String("password", "", "SmartThings password")
	flag.Parse()
	client := smartthings.Client{
		Domain:   *domain,
		Username: *username,
		Password: *password,
	}

	switch flag.Arg(0) {
	case "mode":
		mode, err := client.Mode()
		if err != nil {
			glog.Fatal(err)
		}
		fmt.Println(mode)

	default:
		flag.PrintDefaults()
	}
}
