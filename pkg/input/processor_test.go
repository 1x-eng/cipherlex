package input

import (
	"testing"

	"github.com/1x-eng/cipherlex/pkg/config"
	"github.com/stretchr/testify/assert"
)

// TestLoadInputs_Success checks if the input is loaded correctly.
func TestLoadInputs_Success(t *testing.T) {
	processor := NewProcessor(config.InputConfig{
		MinLineLength: 2,
		MaxLineLength: 50,
		MaxLineCount:  100,
	})

	lines, err := processor.LoadInputs("../../test_data/input.txt")
	assert.NoError(t, err)
	assert.NotEmpty(t, lines, "Input file should not be empty")
}

// TestLoadInputs_LineLengthConstraints checks the line length constraints.
func TestLoadInputs_LineLengthConstraints(t *testing.T) {
	processor := NewProcessor(config.InputConfig{
		MinLineLength: 1,
		MaxLineLength: 50,
		MaxLineCount:  100,
	})

	lines, err := processor.LoadInputs("../../test_data/input.txt")
	assert.NoError(t, err)

	for _, line := range lines {
		assert.GreaterOrEqual(t, len(line), 1, "Line should meet the minimum length constraint")
		assert.LessOrEqual(t, len(line), 50, "Line should meet the maximum length constraint")
	}
}

// TestLoadInputs_MaxLineCount checks for maximum line count constraint.
func TestLoadInputs_MaxLineCount(t *testing.T) {
	maxLines := 2
	processor := NewProcessor(config.InputConfig{
		MinLineLength: 1,
		MaxLineLength: 5,
		MaxLineCount:  maxLines,
	})

	lines, err := processor.LoadInputs("../../test_data/input_invalid.txt")
	assert.NoError(t, err)
	assert.Len(t, lines, maxLines, "The number of loaded lines should not exceed the maximum count")
}
