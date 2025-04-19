package worddictionary

/*
Create a class called WordDictionary that has these method
* AddWord: adds word to the dictionary
* Search: search for word inside the dictionary but it should also support "." wildcard

Example:
words in dictionary: ["apple", "car", "bus"]
search "car" -> True
search "bu" -> False
search ".us" -> True
search "a" -> False

Thoughts:
If it is only exact match, a hashmap will do
However, since it has to support the "." wildcar, we have to use a different datastructure
For example, trie will do. That data structure allows storing prefixes of the word
However, we need to add a certain field to indicate that the middle node can also act as the leaf node.
*/

type WordNode struct {
	letter      string
	isEndOfWord bool
	// This helps to check all the possible next words
	next map[string]*WordNode
}

type WordDictionary struct {
	root *WordNode
}

func NewWordDictionary() *WordDictionary {
	return &WordDictionary{
		root: &WordNode{
			next: make(map[string]*WordNode),
		},
	}
}

func (w *WordDictionary) AddWord(word string) {
	tempNode := w.root
	for _, letter := range word {
		nextNode, ok := tempNode.next[string(letter)]
		if !ok {
			newNode := &WordNode{
				letter: string(letter),
				next:   make(map[string]*WordNode),
			}
			tempNode.next[string(letter)] = newNode
			nextNode = newNode
		}
		tempNode = nextNode
	}
	tempNode.isEndOfWord = true
}

func (w *WordDictionary) Search(word string) bool {
	if len(word) == 0 {
		return false
	}
	// use a stack
	type stackWord struct {
		node      *WordNode
		nextIndex int
	}
	stack := []*stackWord{}
	stack = append(stack, &stackWord{node: w.root, nextIndex: 0})
	isEndOfWord := false
	for len(stack) > 0 {
		curStack := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if curStack.nextIndex >= len(word) {
			if curStack.node.isEndOfWord {
				isEndOfWord = true
			}
			continue
		}
		nextLetter := string(word[curStack.nextIndex])
		if nextLetter == "." {
			for _, nextNode := range curStack.node.next {
				stack = append(stack, &stackWord{
					node:      nextNode,
					nextIndex: curStack.nextIndex + 1,
				})
			}
		} else {
			nextNode, ok := curStack.node.next[nextLetter]
			if !ok {
				continue
			}
			stack = append(stack, &stackWord{
				node:      nextNode,
				nextIndex: curStack.nextIndex + 1,
			})
		}
	}
	return isEndOfWord
}
