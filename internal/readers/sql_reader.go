package readers

import (
    // in build package
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
    "regexp"

    // github package
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/evertras/bubble-table/table"

    // DIY package
	parse "deta/internal/parseSQL"
	"deta/utils"
)

type sqlLoaderMsg struct{
	titles []string
	rows [][]string
    name string
}

type sqlErrorMsg error

// Struct for SQL model
type ModelReaderSQL struct {
    Name string
	Path string
	Cursor int
	Title *[]string
	Rows *[][]string
	Table table.Model
	Style struct{
		Title lipgloss.Style
		Item lipgloss.Style
		Cancel lipgloss.Style
		FontType lipgloss.Style
	}
}

type SQLReader struct{
    path string
    model *ModelReaderSQL
}

// Constractor 
func NewSQLReader() *SQLReader{
    return &SQLReader{
        model: &ModelReaderSQL{},
    }
}

// Run model
func(r *SQLReader) Run() error{
    p := tea.NewProgram(r.model)
    if _, err := p.Run(); err != nil{
        return err
    }
     
    return nil
}

// Отрисовка модели 
func (m *ModelReaderSQL) View() string{
	return m.Table.View() + "\n↑/↓: navigate • q: quit"
}

func readSQL(path string) tea.Cmd {
    return func() tea.Msg {
        if err := utils.CheckFile(path); err != nil{
            fmt.Println("File don't exists in directory")
            return  sqlErrorMsg(err)
        }

        file, err := os.Open(path)
        if err != nil {
            return csvErrorMsg(err)
        }
        defer file.Close()

        // regular expression for clearing comments
        re := regexp.MustCompile(`(--|#).*`)
        title := make([]string, 0)
        rows := make([][]string, 0)
        var titleDefalt parse.CellDefault

        scanner := bufio.NewScanner(file)
        resultModel := sqlLoaderMsg{}
        var sqlQuery string
        for scanner.Scan() {
            line := scanner.Text()
            line = re.ReplaceAllString(line, "")
            line = strings.TrimSpace(line)
            if strings.Contains(line, ";"){
                sqlQuery += line
                temp := parse.Parse(line)
                
                switch temp.Flag {
                case "t":
                    resultModel.name = temp.Query[0]
                    titleDefalt = parse.ParseTableCell(temp.Query[1])
                    resultModel.titles = titleDefalt.Title
                    sqlQuery = ""
                case "i":
                    if resultModel.name == ""{
                        fmt.Println("Table don't exists")
                        return csvErrorMsg(errors.New("Table don't exists"))
                    }
                    
                }
            } else{
                sqlQuery += line + " "
            }

        }

        if err := scanner.Err(); err != nil {
            return csvErrorMsg(err)
        }	

        return resultModel
    }
}


// Initilization model sql
func(r *SQLReader) Init(path string) error{
    r.path = path
	title := make([]string, 0)
	rows := make([][]string, 0)
	
	r.model = &ModelReaderSQL{
		Path:  path,
		Title: &title,
		Rows:  &rows,
	}
	
	r.model.Style.Item = lipgloss.NewStyle().Background(lipgloss.Color("#5CC0C2"))
	r.model.Style.Title = lipgloss.NewStyle().Background(lipgloss.Color("#019395"))
	r.model.Style.FontType = lipgloss.NewStyle().Foreground(lipgloss.Color("#6196A8"))

	return nil
}

// Initialization model 
func(m *ModelReaderSQL) Init() tea.Cmd{
	return readSQL(m.Path)
}

// Update table
func (m *ModelReaderSQL) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case sqlLoaderMsg:
        m.Title = &msg.titles
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
        switch msg.String(){
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "down":
			var cmd tea.Cmd
			m.Table, cmd = m.Table.Update(msg)
			return m, cmd
        case "left", "right":
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

