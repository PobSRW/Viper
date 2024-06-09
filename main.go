package main

import (
	"fmt"
	"log"
	"obp-viper/config"
)

const ()

func main() {
	config, err := config.LoadConfiguration()
	if err != nil {
		log.Fatalf("could not load configuratiton: %v", err)
	}

	fmt.Println(config.Environment)
}
