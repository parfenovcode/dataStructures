package main

import "fmt"

const alphabetSize = 26

//	Node presents each node in the trie
type Node struct {
	children [alphabetSize]*Node
	isEnd    bool
}

// Trie represents a trie and has a pointer to the root node
type Trie struct {
	root *Node
}

// INitTrie will create a new Trie
func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

// Insert will take in a word and add it to the trie
func (t *Trie) Insert(w string) {

	wordLength := len(w)
	currentNode := t.root

	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}

		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

// Serch node will tale in a word and RETURN true if that word is in ithe Trie
func (t *Trie) Search(w string) bool {

	wordLength := len(w)
	currentNode := t.root

	for i := 0; i < wordLength; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}

	return false
}

func main() {
	myTrie := InitTrie()
	myTrie.Insert("word")
	fmt.Println(myTrie.Search("word"))

}
