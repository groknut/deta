package readers

import (
    "os"
    "io"

    tea "github.com/charmbracelet/bubbletea"
    "github.com/evertras/bubble-table/table"

    "deta/utils"
    parse "deta/internal/parse"
)

type hexVisibleMsg struct {
    rows []table.Row
}
type binaryErrorMsg error

type ModelReaderBinary struct {
    file        *os.File
    fileSize    int64
    rowsPerPage int
    cursorRow   int64
    selectedRow int
    offset      int64
    table       table.Model
    styles      utils.ReaderStyles
}

func BuildRows(offset int64, data []byte, rowsPerPage int) []table.Row {
    rows := make([]table.Row, rowsPerPage)
    for i := 0; i < rowsPerPage; i++ {
        start := i * parse.BytesPerRow
        end := start + parse.BytesPerRow
        var chunk []byte
        if start < len(data) {
            if end > len(data) {
                chunk = data[start:len(data)]
            } else {
                chunk = data[start:end]
            }
        } else {
            chunk = nil
        }

        rowOffset := offset + int64(start)
        hexStr, asciiStr := parse.FormatHexAndASCII(chunk)
        rows[i] = table.NewRow(table.RowData{
            "Offset": parse.FormatOffset(rowOffset),
            "Hex":    hexStr,
            "ASCII":  asciiStr,
        })
    }
    return rows
}

func (m *ModelReaderBinary) totalRows() int64 {
    return parse.TotalRows(m.fileSize)
}

func (m *ModelReaderBinary) updateViewport() {
    m.offset, m.selectedRow = parse.CalcViewport(
        m.cursorRow,
        m.totalRows(),
        int64(m.rowsPerPage),
    )
}

func (m *ModelReaderBinary) loadVisibleRows() tea.Cmd {
    return func() tea.Msg {
        if m.file == nil || m.rowsPerPage == 0 {
            return hexVisibleMsg{rows: nil}
        }

        buf := make([]byte, m.rowsPerPage*parse.BytesPerRow)
        n, err := m.file.ReadAt(buf, m.offset)
        if err != nil && err != io.EOF {
            return binaryErrorMsg(err)
        }

        rows := BuildRows(m.offset, buf[:n], m.rowsPerPage)
        return hexVisibleMsg{rows: rows}
    }
}

func (m *ModelReaderBinary) Init() tea.Cmd {
	return m.loadVisibleRows()
}

func (m *ModelReaderBinary) View() string {
	return m.table.View() + "\n↑↓: row • PgUp/PgDn: page • Home/End • q: quit"
}

func (m *ModelReaderBinary) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    case hexVisibleMsg:
        columns := []table.Column{
            table.NewColumn("Offset", "Offset", 10),
            table.NewColumn("Hex", "Hex", 3*parse.BytesPerRow+2),
            table.NewColumn("ASCII", "ASCII", parse.BytesPerRow),
        }

        t := table.New(columns).
            WithRows(msg.rows).
            Focused(true).
            WithPageSize(m.rowsPerPage + 1).
            HeaderStyle(m.styles.Title).
            HighlightStyle(m.styles.Item)

        if len(msg.rows) > 0 {
            if m.selectedRow >= len(msg.rows) {
                m.selectedRow = len(msg.rows) - 1
            }
        } else {
            m.selectedRow = 0
        }
        t = t.WithHighlightedRow(m.selectedRow)
        m.table = t
        return m, nil

    case binaryErrorMsg:
        if m.file != nil {
            m.file.Close()
        }
        return m, tea.Quit

    case tea.WindowSizeMsg:
        m.table = m.table.WithMaxTotalWidth(msg.Width)
        return m, nil

    case tea.KeyMsg:
        switch msg.String() {
        case "q", "ctrl+c":
            if m.file != nil {
                m.file.Close()
            }
            return m, tea.Quit

        case "up":
            if m.cursorRow > 0 {
                m.cursorRow--
                m.updateViewport()
                return m, m.loadVisibleRows()
            }
        case "down":
            total := m.totalRows()
            if total > 0 && m.cursorRow < total-1 {
                m.cursorRow++
                m.updateViewport()
                return m, m.loadVisibleRows()
            }
        case "pgup":
            m.cursorRow -= int64(m.rowsPerPage)
            if m.cursorRow < 0 {
                m.cursorRow = 0
            }
            m.updateViewport()
            return m, m.loadVisibleRows()
        case "pgdown":
            total := m.totalRows()
            m.cursorRow += int64(m.rowsPerPage)
            if m.cursorRow >= total {
                m.cursorRow = total - 1
                if m.cursorRow < 0 {
                    m.cursorRow = 0
                }
            }
            m.updateViewport()
            return m, m.loadVisibleRows()
        case "home":
            m.cursorRow = 0
            m.updateViewport()
            return m, m.loadVisibleRows()
        case "end":
            m.cursorRow = m.totalRows() - 1
            if m.cursorRow < 0 {
                m.cursorRow = 0
            }
            m.updateViewport()
            return m, m.loadVisibleRows()
        }
    }

    return m, nil
}

type BinaryReader struct {
    path  string
    model *ModelReaderBinary
}

func NewBinaryReader() *BinaryReader {
    return &BinaryReader{
        model: &ModelReaderBinary{},
    }
}

func (r *BinaryReader) Init(path string) error {
    r.path = path
    file, err := os.Open(path)
    if err != nil {
        return err
    }
    stat, err := file.Stat()
    if err != nil {
        file.Close()
        return err
    }

    r.model.file = file
    r.model.fileSize = stat.Size()
    r.model.rowsPerPage = 16
    r.model.cursorRow = 0
    r.model.styles = utils.DefaultStyles()
    return nil
}

func (r *BinaryReader) Run() error {
    p := tea.NewProgram(r.model)
    if _, err := p.Run(); err != nil {
        return err
    }
    if r.model.file != nil {
        r.model.file.Close()
    }
    return nil
}
