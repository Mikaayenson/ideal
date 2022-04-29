package main

import (
	"fmt"

	"github.com/Mikaayenson/ideal/configs"
	"github.com/Mikaayenson/ideal/internal/greetings"
	log "github.com/sirupsen/logrus"
)

func main() {

	configs.SetupLogger()
	config := configs.ReadConfig()
	message, err := greetings.Hello(config.Settings.Username)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(message)
}
