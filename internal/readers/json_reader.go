package readers

import (
    "os"

    tea "github.com/charmbracelet/bubbletea"

    "deta/utils"
    "deta/internal/parse"
)

type jsonTreeMsg struct {
    lines []string
    nodes []*parse.JSONNode
}

type jsonErrorMsg error

type ModelReaderJSON struct {
    file *os.File

    root      *parse.JSONNode
    flatLines []string
    flatNodes []*parse.JSONNode

    cursor int

    styles utils.ReaderStyles
    keyMap utils.KeyMap
}

func (m *ModelReaderJSON) loadTree() tea.Cmd {
    return func() tea.Msg {
        return jsonTreeMsg{
            lines: m.flatLines,
            nodes: m.flatNodes,
        }
    }
}

func (m *ModelReaderJSON) Init() tea.Cmd {
    return m.loadTree()
}

func (m *ModelReaderJSON) View() string {
    // пока заглушка
    return ""
}

func (m *ModelReaderJSON) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    // будет добавлено позже
    return m, nil
}

// обёртка
type JSONReader struct {
    path  string
    model *ModelReaderJSON
}

func NewJSONReader() *JSONReader {
    return &JSONReader{
        model: &ModelReaderJSON{},
    }
}

func (r *JSONReader) Init(path string) error {
    r.path = path
    data, err := os.ReadFile(path)
    if err != nil {
        return err
    }
    root, err := parse.ParseJSON(data)
    if err != nil {
        return err
    }
    r.model.root = root
    r.model.styles = utils.DefaultStyles()
    r.model.keyMap = utils.DefaultKeyMap()
    return nil
}

func (r *JSONReader) Run() error {
    p := tea.NewProgram(r.model)
    if _, err := p.Run(); err != nil {
        return err
    }
    return nil
}
