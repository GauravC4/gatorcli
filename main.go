package main

import (
	"fmt"
	"log"

	"github.com/GauravC4/gatorcli/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("error reading from config file : ", err)
	}

	cfg.SetUser("gaurav")

	fmt.Printf("%+v\n", cfg)
}
