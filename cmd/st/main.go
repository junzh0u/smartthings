package main

import (
	"flag"
	"fmt"

	"github.com/golang/glog"
	"github.com/junzh0u/smartthings"
)

func main() {
	flag.Set("stderrthreshold", "FATAL")
	username := flag.String("username", "", "SmartThings username")
	password := flag.String("password", "", "SmartThings password")
	flag.Parse()
	client := smartthings.Client{
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
		glog.Fatalf("Unsupported command: %s", flag.Arg(0))
	}
}
