package main

import (
	"fmt"

	"github.com/GauravC4/gatorcli/internal/config"
)

func main() {
	cfg := config.Read()
	cfg.SetUser()

	cfg = config.Read()
	fmt.Printf("%+v\n", cfg)
}
