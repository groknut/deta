package parse

// TotalRows возвращает количество строк (по 16 байт) для файла заданного размера.
func TotalRows(fileSize int64) int64 {
    if fileSize == 0 {
        return 0
    }
    return (fileSize + BytesPerRow - 1) / BytesPerRow
}

// CalcViewport вычисляет смещение offset и индекс выделенной строки selectedRow
// так, чтобы окно из rowsPerPage строк было отцентрировано вокруг cursorRow.
func CalcViewport(cursorRow, totalRows, rowsPerPage int64) (offset int64, selectedRow int) {
    half := rowsPerPage / 2
    desiredStart := cursorRow - half
    if desiredStart < 0 {
        desiredStart = 0
    } else if desiredStart+rowsPerPage > totalRows {
        desiredStart = totalRows - rowsPerPage
        if desiredStart < 0 {
            desiredStart = 0
        }
    }
    offset = desiredStart * BytesPerRow
    selectedRow = int(cursorRow - desiredStart)
    return
}
