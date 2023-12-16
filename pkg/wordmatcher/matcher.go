package wordmatcher

import (
	"sort"
	"strings"
	"sync"

	"github.com/1x-eng/cipherlex/pkg/config"
	"github.com/1x-eng/cipherlex/pkg/wordmatcher/trie"
)

// Matcher is a struct that holds the trie and chunk size.
type Matcher struct {
	trie      *trie.Trie
	chunkSize int
}

// creates a new Matcher with the given dictionary and configuration.
func NewMatcher(dict []string, cfg config.AppConfig) *Matcher {
	m := &Matcher{
		trie:      trie.NewTrie(),
		chunkSize: cfg.ChunkSize,
	}
	for _, word := range dict {
		key := generateKey(word)
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
func processChunk(chunk string, t *trie.Trie) map[string]struct{} {
	localMatches := make(map[string]struct{})
	for i := 0; i < len(chunk); i++ {
		for j := i + 1; j <= len(chunk); j++ {
			substr := chunk[i:j]
			key := generateKey(substr)
			if t.Find(key) {
				localMatches[substr] = struct{}{}
			}
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

// utility to generate a key for a given word.
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
	return chunks
}
