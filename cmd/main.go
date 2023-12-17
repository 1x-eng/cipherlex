package main

import (
	"os"

	"github.com/1x-eng/cipherlex/pkg/config"
	"github.com/1x-eng/cipherlex/pkg/orchestrator"
	"github.com/1x-eng/cipherlex/pkg/utils"
)

func main() {
	if len(os.Args) < 3 {
		utils.Log.Fatalf("Usage: %s [PATH TO DICTIONARY FILE] [PATH TO INPUT FILE]", os.Args[0])
	}
	dictionaryFilePath, inputFilePath := os.Args[1], os.Args[2]

	utils.Log.Info("Loading cipherlex configuration")
	appConfig := config.NewAppConfig()

	utils.Log.WithFields(map[string]interface{}{
		"dictionaryPath": dictionaryFilePath,
		"inputPath":      inputFilePath,
	}).Info("Starting processing")

	orchestrator.Processor(dictionaryFilePath, inputFilePath, appConfig)

	utils.Log.Info("Cipherlex completed successfully")
}
