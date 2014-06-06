package main

import (
	"fmt"
	"os"
	"github.com/kubmo/locus/config"
	"github.com/kubmo/locus/twilio"
)

func main() {
	currentDir, _ := os.Getwd()
	config, _ := config.FromFile(fmt.Sprintf("%s/.config", currentDir))
	twilio.SendMessage(config.ApiKey, config.ApiSecret, "Hi Kinsey")
	fmt.Printf("Sent message")
}
