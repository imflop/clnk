package main

import (
	"flag"
	"log"

	"github.com/imflop/clnk/internal/app/clnkserver"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config path", "configs/clnk.yml", "path to config file")
}

func main() {
	flag.Parse()
	config, err := clnkserver.NewConfig(configPath)
	if err != nil {
		log.Fatalln(err)
	}
	if err := clnkserver.Run(config); err != nil {
		log.Fatalln(err)
	}
}
