package cmd

import "testing"

func TestParseStringNum(t *testing.T) {
	t.Logf("test parseStringNum() func")

	tests := []struct {
		stringNum  string
		base, want int
	}{
		{"123", 10, 123},
		{"10101", 2, 21},
		{"ABCDEF", 16, 11259375},
	}

	for _, tt := range tests {
		if got := parseStringNum(tt.stringNum, tt.base); got != tt.want {
			t.Errorf("want = %d, got = %d", tt.want, got)
		}
	}
}
