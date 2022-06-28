package inv

import (
	"testing"
	"fmt"
)

type Trie struct {
	nodes [26]*Trie
	b bool
}

func Constructor() Trie {
	return Trie{}
}

func (this *Trie) Insert(word string) {
	node := this
	for _, b := range word {
		if node.nodes[b-'a'] == nil {
			node.nodes[b-'a'] = &Trie{}
		}
		node = node.nodes[b-'a']
	}
	node.b = true
}

func (this *Trie) Search(word string) bool {
	node := this
	for _, b := range word {
		if node.nodes[b-'a'] == nil {
			return false
		}
		node = node.nodes[b-'a']
	}
	return node.b
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this 
	for _, b := range prefix {
		if node.nodes[b-'a'] == nil {
			return false
		}
		node = node.nodes[b-'a']
	}
	return true
}

func TestSearch(t *testing.T) {
	tr := Constructor() 
	tr.Insert("apple")
	res := tr.Search("apple")
	fmt.Println(res)
}
func TestStartWith(t *testing.T) {
	tr := Constructor()
	tr.Insert("apple")
	res := tr.StartsWith("app")
	fmt.Println(res)
}
