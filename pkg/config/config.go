package config

import (
	"os"
	"strconv"
)

// AppConfig holds all application-wide configuration settings.
type AppConfig struct {
	DictionaryConfig
	InputConfig
	ChunkSize int
}

// DictionaryConfig holds configuration settings specific to dictionary processing.
type DictionaryConfig struct {
	MinWordLength     int
	MaxWordLength     int
	MaxDictionarySize int
}

// InputConfig holds configuration settings specific to input processing.
type InputConfig struct {
	MinLineLength int
	MaxLineLength int
	MaxLineCount  int
}

// NewAppConfig creates a new AppConfig with settings from environment variables.
func NewAppConfig() AppConfig {
	return AppConfig{
		DictionaryConfig: DictionaryConfig{
			MinWordLength:     getEnvAsInt("MIN_WORD_LENGTH", 2),
			MaxWordLength:     getEnvAsInt("MAX_WORD_LENGTH", 20),
			MaxDictionarySize: getEnvAsInt("MAX_DICTIONARY_SIZE", 100),
		},
		InputConfig: InputConfig{
			MinLineLength: getEnvAsInt("MIN_LINE_LENGTH", 2),
			MaxLineLength: getEnvAsInt("MAX_LINE_LENGTH", 500),
			MaxLineCount:  getEnvAsInt("MAX_LINE_COUNT", 100),
		},
		ChunkSize: getEnvAsInt("CHUNK_SIZE", 1024),
	}
}

// utility to get an environment variable as an integer.
func getEnvAsInt(name string, defaultVal int) int {
	if value, exists := os.LookupEnv(name); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultVal
}
