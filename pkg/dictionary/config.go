package dictionary

import (
	"os"
	"strconv"
)

// settings for the dictionary
type Config struct {
	MinWordLength     int
	MaxWordLength     int
	MaxDictionarySize int
}

// interface for obtaining dictionary configuration.
type Configurator interface {
	GetConfig() Config
}

// EnvConfigurator implements Configurator to provide configuration from environment variables.
type EnvConfigurator struct{}

func NewEnvConfigurator() *EnvConfigurator {
	return &EnvConfigurator{}
}

func (e *EnvConfigurator) GetConfig() Config {
	return Config{
		MinWordLength:     getEnvAsInt("MIN_WORD_LENGTH", 2),
		MaxWordLength:     getEnvAsInt("MAX_WORD_LENGTH", 20),
		MaxDictionarySize: getEnvAsInt("MAX_DICTIONARY_SIZE", 100),
	}
}

func getEnvAsInt(name string, defaultVal int) int {
	if value, exists := os.LookupEnv(name); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultVal
}
