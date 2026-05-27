package parse

// TotalRows возвращает количество строк (по 16 байт) для файла заданного размера.
func TotalRows(fileSize int64) int64 {
    if fileSize == 0 {
        return 0
    }
    return (fileSize + BytesPerRow - 1) / BytesPerRow
}
