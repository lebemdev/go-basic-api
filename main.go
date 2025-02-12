package main

import (
	"fmt"
	"go-basic-api/cmd/api"
	"go-basic-api/config"
	"log"
)

func main() {
	fmt.Println("Main.go")
	cfg, err := config.SetConfig()

	if err != nil {
		log.Fatalf("config file is not loaded properly %v\n", err)
	}

	api.SetupApp(cfg)
}
