package main

import (
	"strings"
	"testing"
)

// Test strings.Contains()
func TestContains(t *testing.T) {
	cases := []struct {
		s      string
		sub    string
		result bool
	}{
		{"hello world", "", true},
		{"hello world", "h", true},
		{"hello world", " ", true},
		{"hello world", "wo", true},
		{"hello world", "H", false},
	}

	for _, c := range cases {
		t.Run(c.s, func(*testing.T) {
			if strings.Contains(c.s, c.sub) != c.result {
				t.Errorf("not expected: %v, %v", c.s, c.sub)
			}
		})
	}
}
