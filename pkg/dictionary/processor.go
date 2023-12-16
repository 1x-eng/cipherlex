package dictionary

import (
	"bufio"
	"os"
	"strings"
)

// interface for loading and filtering words from a dictionary.
type DictionaryProcessor interface {
	LoadDictionary(filePath string) ([]string, error)
	ApplyConstraints(words []string) []string
}

// Processor implements the DictionaryProcessor interface.
type Processor struct {
	config Configurator
}

// NewProcessor creates a new Processor with the given configuration.
func NewProcessor(config Configurator) *Processor {
	return &Processor{
		config: config,
	}
}

// LoadDictionary loads the dictionary from a file.
func (p *Processor) LoadDictionary(filePath string) ([]string, error) {
	words, err := p.readWordsFromFile(filePath)
	if err != nil {
		return nil, err
	}

	return p.ApplyConstraints(words), nil
}

// utility to scan words from given file into a slice.
func scanWords(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var words []string
	for scanner.Scan() {
		words = append(words, strings.TrimSpace(scanner.Text()))
	}
	return words
}

// readWordsFromFile reads words from the given file path.
func (p *Processor) readWordsFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	return scanWords(file), nil
}

// isValidWord is a utility to check if the given word is valid according to the configuration.
func isValidWord(word string, config Config) bool {
	return len(word) >= config.MinWordLength && len(word) <= config.MaxWordLength
}

// filterWords filters the given words according to the configuration.
func filterWords(words []string, config Config) []string {
	var filteredWords []string
	wordSet := make(map[string]struct{})

	for _, word := range words {
		if !isValidWord(word, config) {
			continue
		}
		if _, exists := wordSet[word]; exists {
			continue
		}
		if len(filteredWords) >= config.MaxDictionarySize {
			break
		}
		filteredWords = append(filteredWords, word)
		wordSet[word] = struct{}{}
	}

	return filteredWords
}

// ApplyConstraints applies the constraints to the given words.
func (p *Processor) ApplyConstraints(words []string) []string {
	config := p.config.GetConfig()
	return filterWords(words, config)
}
