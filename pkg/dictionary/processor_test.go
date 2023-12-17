package dictionary

import (
	"testing"

	"github.com/1x-eng/cipherlex/pkg/config"
	"github.com/stretchr/testify/assert"
)

// TestLoadDictionary_Success checks if the dictionary is loaded correctly.
func TestLoadDictionary_Success(t *testing.T) {
	processor := NewProcessor(config.DictionaryConfig{
		MinWordLength:     2,
		MaxWordLength:     5,
		MaxDictionarySize: 100,
	})

	words, err := processor.LoadDictionary("../../test_data/dict.txt")
	assert.NoError(t, err)
	assert.NotEmpty(t, words, "Dictionary should not be empty")
}

// TestLoadDictionary_WordConstraints checks the word length constraints.
func TestLoadDictionary_WordConstraints(t *testing.T) {
	processor := NewProcessor(config.DictionaryConfig{
		MinWordLength:     3,
		MaxWordLength:     40,
		MaxDictionarySize: 100,
	})

	words, err := processor.LoadDictionary("../../test_data/dict.txt")
	assert.NoError(t, err)

	for _, word := range words {
		assert.GreaterOrEqual(t, len(word), 3, "Word should meet the minimum length constraint")
		assert.LessOrEqual(t, len(word), 40, "Word should meet the maximum length constraint")
	}
}

// TestLoadDictionary_DuplicateWords checks for elimination of duplicate words.
func TestLoadDictionary_DuplicateWords(t *testing.T) {
	processor := NewProcessor(config.DictionaryConfig{
		MinWordLength:     2,
		MaxWordLength:     10,
		MaxDictionarySize: 100,
	})

	words, err := processor.LoadDictionary("../../test_data/dict_invalid.txt")
	assert.NoError(t, err)

	wordSet := make(map[string]bool)
	for _, word := range words {
		assert.False(t, wordSet[word], "Duplicate word found: "+word)
		wordSet[word] = true
	}
}

// TestLoadDictionary_MaxSize checks if the maximum size constraint is respected.
func TestLoadDictionary_MaxSize(t *testing.T) {
	maxSize := 3
	processor := NewProcessor(config.DictionaryConfig{
		MinWordLength:     1,
		MaxWordLength:     5,
		MaxDictionarySize: maxSize,
	})

	words, err := processor.LoadDictionary("../../test_data/dict_invalid.txt")
	assert.NoError(t, err)
	assert.Len(t, words, maxSize, "The number of loaded words should not exceed the maximum size")
}
