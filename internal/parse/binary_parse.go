package parse

import (
    "fmt"
)

// FormatOffset возвращает 8-значное HEX-смещение (например "00000010").
func FormatOffset(offset int64) string {
    return fmt.Sprintf("%08X", offset)
}
