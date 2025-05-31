package word_break_139

/**
This solution is time out. Because it uses DFS to search for solutions.
**/
type TrieNode struct {
	Nexts []*TrieNode
	Leaf  bool
}

func (t *TrieNode) Add(word []byte) {
	if len(word) == 0 {
		t.Leaf = true
		return
	}
	nextNode := t.Nexts[word[0]-'a']
	if nextNode == nil {
		nextNode = &TrieNode{
			Nexts: make([]*TrieNode, 26),
			Leaf:  false,
		}
		t.Nexts[word[0]-'a'] = nextNode
	}
	nextNode.Add(word[1:])
}

func (t *TrieNode) NextNode(c byte) (*TrieNode, bool) {
	n := t.Nexts[c-'a']
	return n, n != nil
}

func wordBreakDFS(s string, wordDict []string) bool {
    
	trie := &TrieNode{
		Nexts: make([]*TrieNode, 26),
		Leaf:  false,
	}
	for _, word := range wordDict {
		trie.Add([]byte(word))
	}

	stack := make([]int, 0, len(wordDict))
	prev := trie

	for i := 0; ; {
		if i == len(s) {
			if prev.Leaf {
				return true
			} else if len(stack) == 0 {
				return false
			} else {
				i = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				prev = trie
			}
		}
		ch := s[i]
		next, ok := prev.NextNode(ch)
		if !ok {
			if !prev.Leaf {
				if len(stack) > 0 {
					i = stack[len(stack)-1]
					stack = stack[:len(stack)-1]
					prev = trie
				} else {
					return false
				}
			} else {
				prev = trie
			}
		} else {
			if prev.Leaf {
				stack = append(stack, i)
			}
			i++
			prev = next

		}
	}
	return prev.Leaf
}