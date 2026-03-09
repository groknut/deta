package readers

import (
	"bufio"
	"context"
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
	// "github.com/charmbracelet/bubbletable"

	// "github.com/Evertras/bubble-table/table"
	"github.com/evertras/bubble-table/table"
	tea "github.com/charmbracelet/bubbletea"
)



// Структура для модели
type ModelReaderCSV struct {
	Path string
	Cursor int
	Title *[]string
	Rows *[][]string
	CtxCancel context.CancelFunc
	Ctx context.Context
	Table table.Model
	Style struct{
		Title lipgloss.Style
		Item lipgloss.Style
		Cancel lipgloss.Style
	}
}

//Старт работы модели
func StartReaderCSV(ctx context.Context,ctxCancel context.CancelFunc, pathFile string) error{
	title := make([]string,0)
	rows :=  make([][]string,0)
	model := ModelReaderCSV{Path: pathFile, 
							Ctx: ctx,
							Title: &title,
							Rows: &rows,
							CtxCancel: ctxCancel}
	model.Style.Item = lipgloss.NewStyle().Background(lipgloss.Color("#C40361"))
	model.Style.Title = lipgloss.NewStyle().Background(lipgloss.Color("#8C0286"))

	model.Init()	
	p := tea.NewProgram(model)
						
	select{
	case <-model.Ctx.Done():
		return errors.New("Error reader")
	default:
		go model.ReadCSV()
	}
	if _, err := p.Run(); err != nil {
		return err
	}
	

	return nil
}


// Отрисовка модели 
func (m ModelReaderCSV) View() string{
	return m.Table.View() + "\n\n↑/↓: navigate • q: quit"
}

// Обновление инфлормации в таблицы
func(m ModelReaderCSV) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	switch msg := msg.(type){
	case tea.KeyMsg:

		switch msg.String(){
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "down":
			
				// newTable, cmd := m.Table.Update(msg)
				// m.Table = newTable.(tea.Model)
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


// Инициализируем нашу модель
func(m ModelReaderCSV) Init() tea.Cmd{
	return nil
}

//Чтение csv файла
func(m *ModelReaderCSV) ReadCSV(){
	file, err := os.Open(m.Path)
	if err != nil{
		m.CtxCancel()
	}
	defer file.Close()
	flagTitle := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		line = strings.TrimSpace(line)
		re, err := regexp.Compile(`[,;|\t]`)
		if err != nil{
			m.CtxCancel()
		}
		words := re.Split(line,-1)
		if flagTitle{
			tempTitle := append(*m.Title, words...)
			*m.Title = tempTitle
			flagTitle = false
		}else{
			 
			tempRows := append(*m.Rows, words)
			*m.Rows = tempRows
		}

	}

	if err := scanner.Err(); err != nil{
		m.CtxCancel()
	}

	width :=  len(*m.Title)
	columns := make([]table.Column, width)
	for i, title := range *m.Title{
		maxWidth := width
		for _, row := range *m.Rows{
			if i < len(row) && len(row[i]) > maxWidth{
				maxWidth = len(row[i])
			}
		}
		// columns[i] = table.Column{
		// 	Title: title,
		// 	Width: maxWidth +2,
		// }
		columns[i] = table.NewColumn(title, title, maxWidth+2)


	}

	tableRows := make([]table.Row, len(*m.Rows))
	for _, row := range *m.Rows {
		rowData := table.RowData{}
		for i := range columns{
			if i < len(row){
				rowData[(*m.Title)[i]] = row[i]
			} else{
				rowData[(*m.Title)[i]] = ""
			}
		}
		// tableRows[i] = table.Row{
		// 	Cells: cells,
		// }
		tableRows = append(tableRows, table.NewRow(rowData))
	}
	m.Table = table.New(columns).
    WithRows(tableRows).
    Focused(true).
    WithPageSize(20)

	m.Table = m.Table.HeaderStyle(m.Style.Title)
	m.Table = m.Table.HighlightStyle(m.Style.Item)
// m.Table = table.New(columns).
//     WithRows(tableRows).
//     WithStyles(table.Styles{
//         Header:   m.Style.Title,
//         Selected: m.Style.Item, // или Highlighted: m.Style.Item (зависит от версии)
//     }).
//     Focused(true).
//     WithPageSize(20)
}