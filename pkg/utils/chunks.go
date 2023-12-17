package utils

import "github.com/1x-eng/cipherlex/pkg/config"

type ChunkSizeCalculator struct {
	config config.InputConfig
}

// returns the length of the longest word in the given slice, this is used to determine the chunk size for the trie.
func LongestWordLength(words []string) int {
	maxLength := 0
	for _, word := range words {
		if len(word) > maxLength {
			maxLength = len(word)
		}
	}
	return maxLength
}

// calculates the average line length of the given lines, this is also required to determine the chunk size.
func CalculateAverageLineLength(lines []string) int {
	totalLength := 0
	for _, line := range lines {
		totalLength += len(line)
	}
	if len(lines) == 0 {
		return 0 // Avoid division by zero
	}
	return totalLength / len(lines)
}

// creates a new ChunkSizeCalculator with the given configuration and dynamically accounting for the longest word length, average line length, and min/max chunk size.
func NewChunkSizeCalculator(cfg config.InputConfig) *ChunkSizeCalculator {
	return &ChunkSizeCalculator{config: cfg}
}

// determines the chunk size based on the longest word length and average line length.
func (calc *ChunkSizeCalculator) DetermineChunkSize(longestWordLength int, averageLineLength int) int {
	chunkSize := longestWordLength

	// if the average line length is greater than the chunk size, and the average line length divided by the chunk size adjustment factor is greater than the minimum chunk size, then adjust the chunk size.
	if averageLineLength > chunkSize*calc.config.ChunkSizeAdjustmentFactor &&
		averageLineLength/calc.config.ChunkSizeAdjustmentFactor > calc.config.MinChunkSize {

		proposedSize := averageLineLength / calc.config.ChunkSizeAdjustmentFactor
		if proposedSize < calc.config.MaxChunkSize {
			chunkSize = proposedSize
		} else {
			chunkSize = calc.config.MaxChunkSize
		}
	} else if chunkSize < calc.config.MinChunkSize {
		chunkSize = calc.config.MinChunkSize
	}

	return chunkSize
}
