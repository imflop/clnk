package main

import (
	"flag"
	"log"

	"github.com/imflop/clnk/internal/app/clnk"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config path", "configs/clnk.yml", "path to config file")
}

func main() {
	flag.Parse()
	config, err := clnk.NewConfig(configPath)
	if err != nil {
		log.Fatalln(err)
	}
	if err := clnk.Run(config); err != nil {
		log.Fatalln(err)
	}
}
