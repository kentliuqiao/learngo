package main

import "testing"

func TestSubStr(t *testing.T) {
	tests := []struct {
		s   string
		ans int
	}{
		// normal cases
		{"abcabcbb", 3},
		{"pkewewpkerew", 5},

		// edge cases
		{"", 0},
		{"aaaaa", 1},
		{"abcabcabcd", 4},

		// chinese support
		{"这里是慕课网慕课", 6},
		{"一二三二一", 3},
	}

	for _, tt := range tests {
		if actual := lengthOfNonRepeatingSubStr(tt.s); actual != tt.ans {
			t.Errorf("non repeating sub str length; string: %s; got: %d; expected: %d", tt.s, actual, tt.ans)
		}
	}
}
