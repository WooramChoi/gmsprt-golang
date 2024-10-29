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
	configFile := flag.String("config", "config/config.yaml", "Config file(.yaml)")
	fmt.Printf("Use [%s]", *configFile)

	flag.Parse()

	// TODO default 파일을 사용하지 않고, Config 객체 생성시 세팅(=default 파일이 없을 경우, 객체 세팅값 사용)
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

	server.Run(config)
}
