package main

import (
	"fmt"
	"testing"
)

func TestComp(t *testing.T) {
	tt := []struct {
		a, b   string
		common string
		ok     bool
	}{
		{"abcde", "fghij", "", false},
		{"abcde", "abcue", "abce", true},
	}

	for _, tc := range tt {
		t.Run(fmt.Sprintf("%s sv %s", tc.a, tc.b), func(t *testing.T) {
			c, ok := comp(tc.a, tc.b)
			if tc.ok != ok {
				t.Fatalf("expected ok = %v, got %v", tc.ok, ok)
			}
			if tc.common != c {
				t.Fatalf("extcted common = %s, got %s", tc.common, c)
			}
		})
	}
}
