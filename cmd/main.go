package main

import (
	"log"
	"os"

	"github.com/1x-eng/cipherlex/pkg/config"
	"github.com/1x-eng/cipherlex/pkg/orchestrator"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s [PATH TO DICTIONARY FILE] [PATH TO INPUT FILE]", os.Args[0])
	}
	dictionaryFilePath, inputFilePath := os.Args[1], os.Args[2]
	appConfig := config.NewAppConfig()

	orchestrator.Processor(dictionaryFilePath, inputFilePath, appConfig)
}
