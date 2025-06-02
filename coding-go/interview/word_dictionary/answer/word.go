package answer

// more elegant solution
type WordDictionary struct {
	nodes     map[rune]*WordDictionary
	endOfWord bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		nodes: make(map[rune]*WordDictionary),
	}
}

func (this *WordDictionary) AddWord(word string) {
	for _, letter := range word {
		nextNode, ok := this.nodes[letter]
		if !ok {
			nextNode = &WordDictionary{
				nodes: make(map[rune]*WordDictionary),
			}
			this.nodes[letter] = nextNode
		}
		this = nextNode
	}
	this.endOfWord = true
}

func (this *WordDictionary) Search(word string) bool {
	for i, letter := range word {
		if _, found := this.nodes[letter]; !found {
			if letter != '.' {
				return false
			}
			for _, node := range this.nodes {
				if node.Search(word[i+1:]) {
					return true
				}
			}
			return false
		}
		this = this.nodes[letter]
	}
	return this.endOfWord
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
