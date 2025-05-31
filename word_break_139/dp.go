package word_break_139

import (
	"fmt"
)

/**
This solution uses dynamic programming to solve the problem.
This is an optimization over the previous DP if dictionary lookups for substrings become a bottleneck. Instead of creating substrings and looking them up in a HashSet, you can traverse the Trie.

DP State and Base Case: Same as above (dp array, dp[0] = true).
Trie: Build your Trie from wordDict as you've started.
Transition: Iterate i from 0 to len(s) - 1. dp[i] signifies if s[0...i-1] can be segmented. If dp[i] is true: Start a traversal from the trie_root with characters s[j] where j goes from i to len(s) - 1. Let currentNode = trie_root. For j from i to len(s) - 1: Move currentNode to currentNode.NextNode(s[j]). If at any point the character s[j] doesn't lead to a valid next node in the Trie, break this inner loop (no word starting at s[i] can be formed further). If currentNode becomes non-nil and currentNode.Leaf is true, it means the substring s[i...j] is a word in the dictionary. Since dp[i] is true (meaning s[0...i-1] is segmentable) and s[i...j] is a word, then s[0...j] is segmentable. So, set dp[j+1] = true.
Result: dp[len(s)].

When I thought of this this problem, I didn't think of it in DP's way.
*/

func reverseWord(w []byte) []byte {
	reversed := make([]byte, len(w))
	for i, j := len(w)-1, 0; i >= 0 && j < len(w); i, j = i-1, j+1 {
		reversed[j] = w[i]
	}
	return reversed
}

func wordBreak(s string, wordDict []string) bool {
	trie := &TrieNode{
		Nexts: make([]*TrieNode, 26),
		Leaf:  false,
	}
	for _, word := range wordDict {
		w := reverseWord([]byte(word))
		fmt.Println(string(w))
		trie.Add(w)
	}

	breakResult := make([]bool, len(s)+1)
	breakResult[0] = true
	for i := 1; i <= len(s); i++ {
		curr := trie
		for j := i - 1; j >= 0; j-- {
			next, ok := curr.NextNode(s[j])
			if !ok {
				breakResult[i] = false
				break
			} else {
				if next.Leaf {
					if breakResult[j] {
						breakResult[i] = true
						break
					}
				}

				curr = next
			}
		}
	}
	return breakResult[len(s)]
}
