package wordmatcher

import (
	"sort"
	"strings"
	"sync"

	"github.com/1x-eng/cipherlex/pkg/config"
	"github.com/1x-eng/cipherlex/pkg/utils"
)

// Matcher is a struct that holds the trie and chunk size.
type Matcher struct {
	trie      *utils.Trie
	chunkSize int
	dictWords []string
}

// creates a new Matcher with the given dictionary and configuration.
func NewMatcher(dict []string, cfg config.AppConfig, chunkSize int) *Matcher {
	m := &Matcher{
		trie:      utils.NewTrie(),
		chunkSize: chunkSize,
		dictWords: dict,
	}
	for _, word := range dict {
		key := generateKey(word)

		utils.Log.WithFields(map[string]interface{}{
			"word": word,
			"key":  key,
		}).Debug("Inserting word into trie")

		m.trie.Insert(key)
	}
	return m
}

// finds all matches in the given input string, concurrently and in chunks, returning a map of matches.
func (m *Matcher) FindMatches(input string) map[string]struct{} {
	matches := make(map[string]struct{})
	chunks := splitString(input, m.chunkSize)

	var wg sync.WaitGroup
	matchMutex := &sync.Mutex{} // Mutex for safely updating 'matches'

	for _, chunk := range chunks {
		wg.Add(1)
		go func(c string) {
			defer wg.Done()
			localMatches := processChunk(c, m.trie)
			mergeMatches(matches, localMatches, matchMutex)
		}(chunk)
	}

	wg.Wait()
	return matches
}

// utility to process a chunk of the input string, finding all matches in the given trie.
func processChunk(chunk string, t *utils.Trie) map[string]struct{} {
	localMatches := make(map[string]struct{})
	for i := 0; i < len(chunk); i++ {
		for j := i + 1; j <= len(chunk); j++ {
			substr := chunk[i:j]
			key := generateKey(substr)
			if t.Find(key) {
				localMatches[substr] = struct{}{}
			}

			utils.Log.WithFields(map[string]interface{}{
				"chunk":      chunk,
				"substr":     substr,
				"key":        key,
				"matchFound": t.Find(key),
			}).Debug("Processed chunk")

		}
	}
	return localMatches
}

// utility to merge local matches processed concurrently into global matches, using a mutex for safety.
func mergeMatches(global, local map[string]struct{}, mutex *sync.Mutex) {
	mutex.Lock()
	defer mutex.Unlock()
	for key := range local {
		global[key] = struct{}{}
	}
}

// utility to generate a key for a given word, by sorting letters to account for anagrams.
func generateKey(word string) string {
	chars := strings.Split(word, "")
	sort.Strings(chars)
	return strings.Join(chars, "")
}

// utility to split a string into chunks of the given size.
func splitString(input string, chunkSize int) []string {
	var chunks []string
	for i := 0; i < len(input); i += chunkSize {
		end := i + chunkSize
		if end > len(input) {
			end = len(input)
		}
		chunks = append(chunks, input[i:end])
	}

	utils.Log.WithFields(map[string]interface{}{
		"input":     input,
		"chunkSize": chunkSize,
		"chunks":    chunks,
	}).Debug("Split input string into chunks")

	return chunks
}

// counts the unique occurrences of dictionary words in the matches.
func (m *Matcher) CountUniqueMatches(matches map[string]struct{}) int {
	uniqueWords := make(map[string]struct{})
	for match := range matches {
		for _, word := range m.dictWords {
			if generateKey(word) == generateKey(match) {
				uniqueWords[word] = struct{}{}
			}
		}
	}

	utils.Log.WithFields(map[string]interface{}{
		"matchCount":  len(matches),
		"uniqueCount": len(uniqueWords),
		"uniqueWords": uniqueWords,
		"dictWords":   m.dictWords,
	}).Debug("Counted unique matches")

	return len(uniqueWords)
}
