package main

import (
	"fmt"
	"log"
	"os"

	"github.com/1x-eng/cipherlex/pkg/config"
	"github.com/1x-eng/cipherlex/pkg/dictionary"
	"github.com/1x-eng/cipherlex/pkg/input"
	"github.com/1x-eng/cipherlex/pkg/wordmatcher"
)

func main() {
	if len(os.Args) < 3 {
		log.Fatalf("Usage: %s [PATH TO DICTIONARY FILE] [PATH TO INPUT FILE]", os.Args[0])
	}
	dictionaryFilePath := os.Args[1]
	inputFilePath := os.Args[2]

	appConfig := config.NewAppConfig()

	dictProcessor := dictionary.NewProcessor(appConfig.DictionaryConfig)
	dictWords, err := dictProcessor.LoadDictionary(dictionaryFilePath)
	if err != nil {
		log.Fatalf("Failed to load dictionary: %v", err)
	}

	inputProcessor := input.NewProcessor(appConfig.InputConfig)
	inputs, err := inputProcessor.LoadInputs(inputFilePath)
	if err != nil {
		log.Fatalf("Failed to load input file: %v", err)
	}

	matcher := wordmatcher.NewMatcher(dictWords, appConfig)

	for i, inputLine := range inputs {
		matches := matcher.FindMatches(inputLine)
		uniqueCount := matcher.CountUniqueMatches(matches)
		fmt.Printf("Case #%d: %d\n", i+1, uniqueCount)
	}
}
