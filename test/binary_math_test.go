package test

import (
    "testing"
    parse "deta/internal/parse"
)

func TestTotalRows(t *testing.T) {
    tests := []struct {
        size     int64
        expected int64
    }{
        {0, 0},
        {1, 1},
        {16, 1},
        {17, 2},
        {32, 2},
        {33, 3},
        {1024, 64},
    }
    for _, tt := range tests {
        got := parse.TotalRows(tt.size)
        if got != tt.expected {
            t.Errorf("size %d: got %d, want %d", tt.size, got, tt.expected)
        }
    }
}
