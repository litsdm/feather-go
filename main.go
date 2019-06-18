package main

import (
	"fmt"
	"log"

	"feather.com/internal"
)

func main() {
	configuration, err := config.New()
	if err != nil {
		log.Panicln("Configuration error", err)
	}
	fmt.Println(configuration.Constants.PORT)
}
