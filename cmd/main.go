package main

import (
	"flag"
	"fmt"
	"os"

	"gmsprt-golang/internal/server"

	"gopkg.in/yaml.v3"
)

func main() {

	config := &server.Config{}
	// configFile := flag.String("config", "config/config.yaml", "Config file(.yaml)")
	configFile := flag.String("config", "", "Config file(.yaml)")

	//
	flag.Parse()

	//
	if *configFile != "" {
		fmt.Printf("Use config file: [%s]", *configFile)
		buf, err := os.ReadFile(*configFile)
		if err != nil {
			fmt.Print(err.Error())
			panic(err)
		}

		err = yaml.Unmarshal(buf, config)
		if err != nil {
			fmt.Print(err.Error())
			panic(err)
		}
	} else {
		fmt.Println("Use default value")
		config.Server.Port = 9000
		config.Database.Type = "sqlite"
	}
	fmt.Printf("Config Detail:\n%v\n", config)

	server.Run(config)
}
