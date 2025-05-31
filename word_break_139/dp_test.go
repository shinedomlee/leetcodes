package word_break_139

import "testing"

func Test_wordBreak(t *testing.T) {
	if wordBreak("leetcode", []string{"leet", "code"}) != true {
		t.Errorf("Expected true, got false")
	}
}
