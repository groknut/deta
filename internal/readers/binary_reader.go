package readers

import (
    "os"
    // "io"

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
    return nil
}

func (m *ModelReaderBinary) View() string {
    return ""
}

func (m *ModelReaderBinary) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
