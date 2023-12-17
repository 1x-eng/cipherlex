package utils

// Node represents a node in the Trie.
type Node struct {
	Children map[rune]*Node
	IsWord   bool
}

// Trie represents the Trie data structure.
type Trie struct {
	Root *Node
}

// NewTrie creates a new Trie instance.
func NewTrie() *Trie {
	return &Trie{Root: &Node{Children: make(map[rune]*Node)}}
}

// Insert inserts a word into the Trie.
func (t *Trie) Insert(word string) {
	node := t.Root
	for _, r := range word {
		if child, ok := node.Children[r]; ok {
			node = child
		} else {
			child := &Node{Children: make(map[rune]*Node)}
			node.Children[r] = child
			node = child
		}
	}
	node.IsWord = true
}

// Find checks if a word is in the Trie.
func (t *Trie) Find(word string) bool {
	node := t.Root
	for _, r := range word {
		if child, ok := node.Children[r]; ok {
			node = child
		} else {
			return false
		}
	}
	return node.IsWord
}
