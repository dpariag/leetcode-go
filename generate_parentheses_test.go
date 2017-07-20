// Leetcode: https://leetcode.com/problems/generate-parentheses/#/description
// Given n, write a function to generate all combinations of n pairs of well-formed parentheses.
// Example: Given n = 3, return: ["((()))","(()())","(())()","()(())","()()()"]

// Brute Force: Generate all 2^n combos of n parens, then check which are balanced. O(n*2^n)
// Better: Recursively generate balanced parens (i.e., only add ')' if there is a '(' to match).

package leetcode

import (
	"fmt"
	"sort"
	"strings"
	"testing"
)

// Accepted. 15ms. Beats 80.95% of submissions, ties 14.29% of submissions
func generateParenthesis(n int) []string {
	generated := []string{}
	if n > 0 {
		s := string("")
		generated = generateParens(n, n, s, generated)
	}
	return generated
}

func generateParens(remaining_open int, remaining_closed int, current string, generated []string) []string {
	if remaining_open == 0 && remaining_closed == 0 {
		generated = append(generated, current)
		return generated
	}

	if remaining_open > 0 {
		generated = generateParens(remaining_open-1, remaining_closed, strings.Join([]string{current, "("}, ""), generated)
	}

	if remaining_closed > remaining_open {
		generated = generateParens(remaining_open, remaining_closed-1, strings.Join([]string{current, ")"}, ""), generated)
	}
	return generated
}

func testParens(t *testing.T, n int, expected []string) {
	p := generateParenthesis(n)
	sort.Strings(p)
	sort.Strings(expected)
	if len(p) == len(expected) {
		// Match against expected strings
		for i := 0; i < len(expected); i++ {
			if p[i] != expected[i] {
				t.Errorf("Mismatch. Expected '%s', but got '%s'", expected[i], p[i])
			}
		}
	} else {
		t.Errorf("Generated %d strings, expected %d\n", len(p), len(expected))
		fmt.Println(p)
		fmt.Println(expected)
	}
}

func TestGenerateParens(t *testing.T) {
	testParens(t, 0, []string{})
	testParens(t, 1, []string{"()"})
	testParens(t, 2, []string{"(())", "()()"})
}
