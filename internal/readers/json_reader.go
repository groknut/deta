package readers

import (
    "os"
    "strings"

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

func (m *ModelReaderJSON) rebuildFlat() {
    m.flatNodes = parse.FlattenTree(m.root)
    m.flatLines = make([]string, len(m.flatNodes))
    for i, node := range m.flatNodes {
        m.flatLines[i] = parse.FormatJSONNode(node)
    }
    // коррекция курсора
    if len(m.flatNodes) == 0 {
        m.cursor = 0
        return
    }
    if m.cursor >= len(m.flatNodes) {
        m.cursor = len(m.flatNodes) - 1
    }
    if m.cursor < 0 {
        m.cursor = 0
    }
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
    var sb strings.Builder
    for i, line := range m.flatLines {
        if i == m.cursor {
            sb.WriteString(m.styles.Item.Render(line))
        } else {
            sb.WriteString(line)
        }
        sb.WriteString("\n")
    }
    return sb.String() + "\n↑↓: move • ←→/hl: collapse/expand • q: quit"
}

func (m *ModelReaderJSON) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {

    case jsonTreeMsg:
        m.flatLines = msg.lines
        m.flatNodes = msg.nodes
        return m, nil

    case jsonErrorMsg:
        if m.file != nil {
            m.file.Close()
        }
        return m, tea.Quit

    case tea.WindowSizeMsg:
        return m, nil

    case tea.KeyMsg:

    	action := m.keyMap.Lookup(msg)
        switch action {
	        case utils.ActionQuit:
	            if m.file != nil {
	                m.file.Close()
	            }
	            return m, tea.Quit
	        case utils.ActionUp:
	            if m.cursor > 0 {
	                m.cursor--
	                return m, nil
	            }
	        case utils.ActionDown:
	            if m.cursor < len(m.flatNodes)-1 {
	                m.cursor++
	                return m, nil
	            }
        }

        // дополнительные клавиши для дерева (раскрытие/сворачивание)
        if m.cursor < len(m.flatNodes) {
            node := m.flatNodes[m.cursor]
            switch action {
	            case utils.ActionRight:
	                if len(node.Children) > 0 && !node.Expanded {
	                    node.Expanded = true
	                    m.rebuildFlat()
	                    return m, m.loadTree()
	                }
	            case utils.ActionLeft:
	                if len(node.Children) > 0 && node.Expanded {
	                    node.Expanded = false
	                    m.rebuildFlat()
	                    return m, m.loadTree()
	                }
	            }
        }
    }
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
    r.model.rebuildFlat()
    return nil
}

func (r *JSONReader) Run() error {
    p := tea.NewProgram(r.model)
    if _, err := p.Run(); err != nil {
        return err
    }
    return nil
}
