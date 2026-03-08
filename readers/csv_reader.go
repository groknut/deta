package readers

import (
	"bufio"
	"context"
	"errors"
	"os"
	"regexp"
	"strings"

	"github.com/charmbracelet/lipgloss"
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
	Style struct{
		title lipgloss.Style
		item lipgloss.Style
		cancel lipgloss.Style
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
	model.Style.item = lipgloss.NewStyle().Background(lipgloss.Color("#C40361"))
	model.Style.title = lipgloss.NewStyle().Background(lipgloss.Color("#8C0286"))

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
	s := m.Style.title.Render(strings.Join(*m.Title, "\t"))+"\n"
	for i, row := range *m.Rows{
		cursor := " "
		if m.Cursor == i{
			cursor = ">"
		}
		s += cursor + strings.Join(row, "\t")+"\n"
	}
	s += "q - exit"
	return s
}

// Обновление инфлормации в таблицы
func(m ModelReaderCSV) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	switch msg := msg.(type){
	case tea.KeyMsg:

		switch msg.String(){
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.Cursor > 0{
				m.Cursor--

			}
		case "down":
			if m.Cursor < len(*m.Rows)-1{
				m.Cursor++
			}
		}
	}

	return m, nil
}

// Инициализируем нашу модель
func(m ModelReaderCSV) Init() tea.Cmd{
	return nil
}

//Чтение csv файла
func(m ModelReaderCSV) ReadCSV(){
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



}
