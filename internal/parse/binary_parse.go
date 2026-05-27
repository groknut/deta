package parse

import (
    "fmt"
    "strings"
)

// BytesPerRow — количество байт в одной строке HEX-дампа.
const BytesPerRow = 16

// FormatOffset возвращает 8-значное HEX-смещение (например "00000010").
func FormatOffset(offset int64) string {
    return fmt.Sprintf("%08X", offset)
}

// FormatHexAndASCII преобразует срез байт в HEX-строку и ASCII-представление.
// Если длина chunk меньше BytesPerRow, недостающие байты заменяются пробелами.
func FormatHexAndASCII(chunk []byte) (hexStr, asciiStr string) {
    hexParts := make([]string, BytesPerRow)
    ascii := make([]byte, BytesPerRow)

    for i := 0; i < BytesPerRow; i++ {
        if i < len(chunk) {
            b := chunk[i]
            hexParts[i] = fmt.Sprintf("%02X", b)
            if b >= 32 && b <= 126 {
                ascii[i] = b
            } else {
                ascii[i] = '.'
            }
        } else {
            hexParts[i] = "  "
            ascii[i] = ' '
        }
    }
    hexStr = strings.Join(hexParts, " ")
    asciiStr = string(ascii)
    return
}
