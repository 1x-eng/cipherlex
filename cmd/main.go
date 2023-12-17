package main

import (
	"flag"
	"os"

	"github.com/1x-eng/cipherlex/pkg/config"
	"github.com/1x-eng/cipherlex/pkg/orchestrator"
	"github.com/1x-eng/cipherlex/pkg/utils"
)

func main() {
	dictionaryFilePath := flag.String("dictionary", "", "Path to dictionary file")
	inputFilePath := flag.String("input", "", "Path to input file")

	flag.Parse()

	if *dictionaryFilePath == "" || *inputFilePath == "" {
		utils.Log.Fatalf("Usage: %s --dictionary [PATH TO DICTIONARY FILE] --input [PATH TO INPUT FILE]", os.Args[0])
	}

	utils.Log.Info("Loading cipherlex configuration")
	appConfig := config.NewAppConfig()

	utils.Log.WithFields(map[string]interface{}{
		"dictionaryPath": dictionaryFilePath,
		"inputPath":      inputFilePath,
	}).Info("Starting processing")

	orchestrator.Processor(*dictionaryFilePath, *inputFilePath, appConfig)

	utils.Log.Info("Cipherlex completed successfully")
}
