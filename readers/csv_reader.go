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
type modelReaderCSV struct {
	path string
	cursor int
	title *[]string
	rows *[][]string
	ctx	context.Context
}

//Старт работы модели
func StartReaderCSV(ctx context.Context, pathFile string) error{
	title := make([]string,0)
	rows :=  make([][]string,0)
	model := modelReaderCSV{path: pathFile, 
							ctx: ctx,
							title: &title,
							rows: &rows}


	model.Init()						
	select{
	case <-model.ctx.Done():
		return errors.New("Error reader")
	default:

	}
	

	return nil
}


// Отрисовка модели 
func (m modelReaderCSV) View() string{
	s := strings.Join(*m.title, "\t")
	for i, row := range *m.rows{
		cursor := " "
		if m.cursor == i{
			cursor = ">"
		}
		s += cursor + strings.Join(row, "\t")
	}
	s += "q - exit"
	return s
}

// Обновление инфлормации в таблицы
func(m modelReaderCSV) Update(msg tea.Msg) (tea.Model, tea.Cmd){
	switch msg := msg.(type){
	case tea.KeyMsg:

		switch msg.String(){
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			if m.cursor > 0{
				m.cursor--
			}
		case "down":
			if m.cursor < len(*m.rows)-1{
				m.cursor++
			}
		}
	}

	return m, nil
}

// Инициализируем нашу модель
func(m modelReaderCSV) Init() tea.Cmd{
	return nil
}

//Чтение csv файла
func(m modelReaderCSV) readCSV(){
	file, err := os.Open(m.path)
	if err != nil{
		m.ctx.Deadline()
	}
	defer file.Close()
	flagTitle := true

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		line = strings.TrimSpace(line)
		re, err := regexp.Compile(`[,;|\t]`)
		if err != nil{
			m.ctx.Deadline()
		}
		words := re.Split(line,-1)
		if flagTitle{
			tempTitle := append(*m.title, words...)
			*m.title = tempTitle
			flagTitle = false
		}else{
			 
			tempRows := append(*m.rows, words)
			*m.rows = tempRows
		}

	}

	if err := scanner.Err(); err != nil{
		m.ctx.Deadline()
	}



}
