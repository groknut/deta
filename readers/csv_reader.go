package readers

import (
	"bufio"
	"context"
	"errors"
	"os"
	"regexp"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

// Структура для модели
type ModelReaderCSV struct {
	Path string
	Cursor int
	Title *[]string
	Rows *[][]string
	Ctx	context.Context
}

//Старт работы модели
func StartReaderCSV(ctx context.Context, pathFile string) error{
	title := make([]string,0)
	rows :=  make([][]string,0)
	model := ModelReaderCSV{Path: pathFile, 
							Ctx: ctx,
							Title: &title,
							Rows: &rows}


	model.Init()						
	select{
	case <-model.Ctx.Done():
		return errors.New("Error reader")
	default:

	}
	

	return nil
}


// Отрисовка модели 
func (m ModelReaderCSV) View() string{
	s := strings.Join(*m.Title, "\t")
	for i, row := range *m.Rows{
		cursor := " "
		if m.Cursor == i{
			cursor = ">"
		}
		s += cursor + strings.Join(row, "\t")
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
func(m ModelReaderCSV) readCSV(){
	file, err := os.Open(m.Path)
	if err != nil{
		m.Ctx.Deadline()
	}
	defer file.Close()
	flagTitle := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		line = strings.TrimSpace(line)
		re, err := regexp.Compile(`[,;|\t]`)
		if err != nil{
			m.Ctx.Deadline()
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
		m.Ctx.Deadline()
	}



}
