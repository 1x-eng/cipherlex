package input

import (
	"bufio"
	"os"
	"strings"

	"github.com/1x-eng/cipherlex/pkg/config"
	"github.com/1x-eng/cipherlex/pkg/utils"
)

// interface for loading and validating input strings.
type InputProcessor interface {
	LoadInputs(filePath string) ([]string, error)
}

// Processor implements the InputProcessor interface.
type Processor struct {
	config config.InputConfig
}

// creates a new input Processor with the given configuration.
func NewProcessor(config config.InputConfig) *Processor {
	return &Processor{
		config: config,
	}
}

// LoadInputs loads and validates input strings from a file.
func (p *Processor) LoadInputs(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		utils.Log.WithError(err).Error("Failed to open input file")
		return nil, err
	}
	defer file.Close()

	return p.scanAndFilterInputs(file), nil
}

// utility to checks if an input line is valid according to the configuration.
func (p *Processor) isValidInput(input string) bool {
	length := len(input)
	isValid := length >= p.config.MinLineLength && length <= p.config.MaxLineLength

	utils.Log.WithFields(map[string]interface{}{
		"input":   input,
		"isValid": isValid,
	}).Debug("Validating input line from input file against configuration")

	return isValid
}

// scanAndFilterInputs scans and filters input lines from a file.
func (p *Processor) scanAndFilterInputs(file *os.File) []string {
	scanner := bufio.NewScanner(file)
	var inputs []string
	for scanner.Scan() {
		input := strings.TrimSpace(scanner.Text())
		if p.isValidInput(input) {
			inputs = append(inputs, input)
			if len(inputs) >= p.config.MaxLineCount {
				utils.Log.WithField("maxLineCount", p.config.MaxLineCount).Warn("Reached max line count, will not process any more lines from input file")
				break
			}
		}
	}
	return inputs
}
