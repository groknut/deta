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

func TestFormatHexAndASCII(t *testing.T) {
    // Полный чанк 16 байт
    full := []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
        0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F}
    hex, ascii := parse.FormatHexAndASCII(full)
    expectedHex := "00 01 02 03 04 05 06 07 08 09 0A 0B 0C 0D 0E 0F"
    expectedASCII := "................" // все непечатные
    if hex != expectedHex || ascii != expectedASCII {
        t.Errorf("full chunk: hex=%q ascii=%q", hex, ascii)
    }

    // Частичный чанк (10 байт)
    partial := []byte{0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x4A}
    hex2, ascii2 := parse.FormatHexAndASCII(partial)
    // Ожидаем 10 байт данных и 6 пробелов в hex
    expectedHex2 := "41 42 43 44 45 46 47 48 49 4A                         "
    // Ожидаем "ABCDEFGHIJ" + 6 пробелов
    expectedASCII2 := "ABCDEFGHIJ      "
    if hex2 != expectedHex2 || ascii2 != expectedASCII2 {
        t.Errorf("partial chunk: hex=%q ascii=%q", hex2, ascii2)
    }

    // Пустой чанк
    emptyHex, emptyASCII := parse.FormatHexAndASCII(nil)
    expectedEmptyHex := "                                                 "
    expectedEmptyASCII := "                "
    if emptyHex != expectedEmptyHex || emptyASCII != expectedEmptyASCII {
        t.Errorf("empty chunk: hex=%q (len=%d) ascii=%q", emptyHex, len(emptyHex), emptyASCII)
    }
}
