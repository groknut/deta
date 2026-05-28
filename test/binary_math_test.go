package test

import (
    "testing"
    parse "github.com/groknut/deta/internal/parse"
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

func TestCalcViewport(t *testing.T) {
    const totalRows = 50
    const rowsPerPage = 16

    tests := []struct {
        cursor         int64
        expectedOffset int64
        expectedSel    int
    }{
        {0, 0, 0},
        {7, 0, 7},
        {8, 0, 8},
        {25, 272, 8},
        {41, 528, 8},
        {49, 544, 15},
    }
    for _, tt := range tests {
        offset, sel := parse.CalcViewport(tt.cursor, totalRows, rowsPerPage)
        if offset != tt.expectedOffset || sel != tt.expectedSel {
            t.Errorf("cursor=%d: offset=%d sel=%d, want offset=%d sel=%d",
                tt.cursor, offset, sel, tt.expectedOffset, tt.expectedSel)
        }
    }
}
