package trie

type Node struct {
	nextMap map[rune]*Node
	letter  rune
    endOfWord bool
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: &Node{
            nextMap: make(map[rune]*Node),
        },
	}
}

func (this *Trie) Insert(word string) {
    node := this.root
	for _, letter := range word {
		nextNode, ok := node.nextMap[letter]
        if !ok {
            nextNode = &Node{
                nextMap: make(map[rune]*Node),
            }
            node.nextMap[letter] = nextNode
        }
        node = nextNode
	}
    node.endOfWord = true
}

func (this *Trie) Search(word string) bool {
	node := this.root
    for _, letter := range word {
        nextNode, ok := node.nextMap[letter]
        if !ok {
            return false
        }
        node = nextNode
    }
    return node.endOfWord
}

func (this *Trie) StartsWith(prefix string) bool {
	node := this.root
    for _, letter := range prefix {
        nextNode, ok := node.nextMap[letter]
        if !ok {
            return false
        }
        node = nextNode
    }
    return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */