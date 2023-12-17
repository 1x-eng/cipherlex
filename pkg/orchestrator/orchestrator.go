package orchestrator

import (
	"fmt"
	"log"

	"github.com/1x-eng/cipherlex/pkg/config"
	"github.com/1x-eng/cipherlex/pkg/dictionary"
	"github.com/1x-eng/cipherlex/pkg/input"
	"github.com/1x-eng/cipherlex/pkg/utils"
	"github.com/1x-eng/cipherlex/pkg/wordmatcher"
)

// Processor is the main entrypoint for the application, it loads and processes the dictionary and input files and then finds matches.
func Processor(dictPath, inputPath string, cfg config.AppConfig) {
	dictWords := loadAndProcessDictionary(dictPath, cfg.DictionaryConfig)
	inputLines := loadAndProcessInput(inputPath, cfg.InputConfig)

	chunkSize := determineChunkSize(dictWords, inputLines, cfg.InputConfig)
	processMatches(inputLines, dictWords, chunkSize, cfg)
}

// loads and processes the dictionary file.
func loadAndProcessDictionary(dictPath string, dictConfig config.DictionaryConfig) []string {
	dictProcessor := dictionary.NewProcessor(dictConfig)
	dictWords, err := dictProcessor.LoadDictionary(dictPath)
	if err != nil {
		log.Fatalf("Failed to load dictionary: %v", err)
	}
	return dictWords
}

// loads and processes the input file.
func loadAndProcessInput(inputPath string, inputConfig config.InputConfig) []string {
	inputProcessor := input.NewProcessor(inputConfig)
	inputLines, err := inputProcessor.LoadInputs(inputPath)
	if err != nil {
		log.Fatalf("Failed to load input file: %v", err)
	}
	return inputLines
}

// dynamically determines the chunk size to use for processing the input file.
func determineChunkSize(dictWords, inputLines []string, inputConfig config.InputConfig) int {
	longestWordLength := utils.LongestWordLength(dictWords)
	averageLineLength := utils.CalculateAverageLineLength(inputLines)
	return utils.NewChunkSizeCalculator(inputConfig).DetermineChunkSize(longestWordLength, averageLineLength)
}

// processes the input lines and finds matches.
func processMatches(inputLines, dictWords []string, chunkSize int, cfg config.AppConfig) {
	matcher := wordmatcher.NewMatcher(dictWords, cfg, chunkSize)
	for i, line := range inputLines {
		matches := matcher.FindMatches(line)
		uniqueCount := matcher.CountUniqueMatches(matches)
		fmt.Printf("Case #%d: %d\n", i+1, uniqueCount)
	}
}
