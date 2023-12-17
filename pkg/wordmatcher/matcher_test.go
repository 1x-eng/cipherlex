package wordmatcher

import (
	"testing"

	"github.com/1x-eng/cipherlex/pkg/config"
	"github.com/stretchr/testify/assert"
)

func TestMatcher_FindMatches(t *testing.T) {
	dict := []string{"axpaj", "apxaj", "dnrbt", "pjxdn", "abd"}
	expectedMatches := []string{"aapxj", "dnrbt", "pxjdn"}
	chunkSize := 10

	matcher := NewMatcher(dict, config.AppConfig{}, chunkSize)
	input := "aapxjdnrbtvldptfzbbdbbzxtndrvjblnzjfpvhdhhpxjdnrbt"

	matches := matcher.FindMatches(input)

	for key := range matches {
		assert.Contains(t, expectedMatches, key, "Expected matches should contain word")
	}
}

func TestMatcher_CountUniqueMatches(t *testing.T) {
	dict := []string{"axpaj", "apxaj", "dnrbt", "pjxdn", "abd"}
	chunkSize := 10

	matcher := NewMatcher(dict, config.AppConfig{}, chunkSize)
	input := "aapxjdnrbtvldptfzbbdbbzxtndrvjblnzjfpvhdhhpxjdnrbt"

	matches := matcher.FindMatches(input)
	uniqueCount := matcher.CountUniqueMatches(matches)

	assert.Equal(t, 4, uniqueCount, "Unique count should be 4")
}
