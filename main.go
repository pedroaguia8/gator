package main

import (
	"fmt"
	"log"

	"github.com/pedroaguia8/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Failed to read initial config: %v", err)
	}

	username := "pedroaguia8"
	err = cfg.SetUser(username)
	if err != nil {
		log.Fatalf("Failed to set user and write config: %v", err)
	}

	updatedCfg, err := config.Read()
	if err != nil {
		log.Fatalf("Failed to read updated config: %v", err)
	}

	fmt.Println(updatedCfg)
}
