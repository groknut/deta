package readers

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/evertras/bubble-table/table"

    "github.com/groknut/deta/utils"
)

type csvLoadedMsg struct {
    title []string
    rows  [][]string
}

type csvErrorMsg error

// Структура для модели
type ModelReaderCSV struct {
	Path string
	Cursor int
	Title *[]string
	Rows *[][]string
	Table table.Model
	Style utils.ReaderStyles
	keyMap utils.KeyMap
}

type CSVReader struct{
    path string
    model *ModelReaderCSV
}

func NewCSVReader() *CSVReader{
    return &CSVReader{
        model: &ModelReaderCSV{},
    }
}

//Запуск модели
func(r *CSVReader) Run() error{
    p := tea.NewProgram(r.model)
    if _,err := p.Run(); err != nil{
        return err
    }

    return nil
}


// Отрисовка модели
func (m *ModelReaderCSV) View() string{
	return m.Table.View() + "\n↑/↓: navigate • q: quit"
}


// Чтение файла
func readCSV(path string) tea.Cmd {
    return func() tea.Msg {
        if err := utils.CheckFile(path); err != nil{
            fmt.Println("File don't exists in directory")
            return  csvErrorMsg(err)
        }
        file, err := os.Open(path)
        if err != nil {
            return csvErrorMsg(err)
        }
        defer file.Close()

        title := make([]string, 0)
        rows := make([][]string, 0)
        flagTitle := true

        scanner := bufio.NewScanner(file)
        for scanner.Scan() {
            line := scanner.Text()
            line = strings.TrimSpace(line)
            re, err := regexp.Compile(`[,;|\t]`)
            if err != nil {
                return csvErrorMsg(err)
            }
            words := re.Split(line, -1)
            if flagTitle {
                title = append(title, words...)
                flagTitle = false
            } else {
                rows = append(rows, words)
            }
        }

        if err := scanner.Err(); err != nil {
            return csvErrorMsg(err)
        }

        return csvLoadedMsg{title: title, rows: rows}
    }
}

// Инициализация интерфейса
func(r *CSVReader) Init(path string) error{
    r.path = path
	title := make([]string, 0)
	rows := make([][]string, 0)

	r.model = &ModelReaderCSV{
		Path:  path,
		Title: &title,
		Rows:  &rows,
	}

	r.model.Style = utils.DefaultStyles()
	r.model.keyMap = utils.DefaultKeyMap()

	return nil
}

// Инициализируем нашу модель
func(m *ModelReaderCSV) Init() tea.Cmd{

	return readCSV(m.Path)
}

// Обновление таблицы
func (m *ModelReaderCSV) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case csvLoadedMsg:
        m.Title = &msg.title
        m.Rows = &msg.rows
        width := len(*m.Title)
        columns := make([]table.Column, width)
        for i, titleStr := range *m.Title {
            maxWidth := len(titleStr)
            for _, row := range *m.Rows {
                if i < len(row) && len(row[i]) > maxWidth {
                    maxWidth = len(row[i])
                }
            }
            columns[i] = table.NewColumn(titleStr, titleStr, maxWidth+2)
        }

        tableRows := make([]table.Row, 0, len(*m.Rows))
        for _, row := range *m.Rows {
            rowData := table.RowData{}
            for i := range columns {
                if i < len(row) {
                    rowData[(*m.Title)[i]] = row[i]
                } else {
                    rowData[(*m.Title)[i]] = ""
                }
            }
            tableRows = append(tableRows, table.NewRow(rowData))
        }

        m.Table = table.New(columns).
            WithRows(tableRows).
            Focused(true).
            WithPageSize(20).
            HeaderStyle(m.Style.Title).
            HighlightStyle(m.Style.Item)

        return m, nil

    case csvErrorMsg:
        return m, tea.Quit

    case tea.KeyMsg:
        action := m.keyMap.Lookup(msg)
        switch action {
        	case utils.ActionQuit:
        		return m, tea.Quit
        	case utils.ActionUp, utils.ActionDown:
        		var cmd tea.Cmd
        		m.Table, cmd = m.Table.Update(msg)
        		return m, cmd
        	case utils.ActionLeft, utils.ActionRight:
            	var cmd tea.Cmd
        	m.Table, cmd = m.Table.Update(msg)
        	return m, cmd
		}

    case tea.WindowSizeMsg:
        m.Table = m.Table.WithMaxTotalWidth(msg.Width)
		m.Table = m.Table.WithPageSize(msg.Height - 5)
    }

    return m, nil
}
