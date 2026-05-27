package test

import (
	parse "deta/internal/parse"
    "testing"
)

func TestFormatOffset(t *testing.T) {
    tests := []struct {
        offset   int64
        expected string
    }{
        {0, "00000000"},
        {16, "00000010"},
        {255, "000000FF"},
        {0xABCDEF, "00ABCDEF"},
    }
    for _, tt := range tests {
        got := parse.FormatOffset(tt.offset)
        if got != tt.expected {
            t.Errorf("FormatOffset(%d) = %s, want %s", tt.offset, got, tt.expected)
        }
    }
}
