package test


import (
	// "fmt"
	// "os/exec"
	// "path/filepath"
	"testing"
	"regexp"
	"os"
	"bufio"
	"fmt"
	"strings"
	parse "github.com/groknut/deta/internal/parseSQL"

)

type sqlLoaderMsg struct{
	titles []string
	rows [][]string
    name string
}

type sqlErrorMsg error

func readFileAddRows(path string){
	file, err := os.Open(path)
	if err != nil {
		return
	}
	defer file.Close()

	re := regexp.MustCompile(`(--|#).*`)
	var titleDefalt parse.CellDefault
	var resultModel sqlLoaderMsg
	var sqlBuilder strings.Builder

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = re.ReplaceAllString(line, "")
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		sqlBuilder.WriteString(line)
		sqlBuilder.WriteString(" ")

		currentSQL := sqlBuilder.String()
		if strings.Contains(currentSQL, ";") {
			semicolonIdx := strings.Index(currentSQL, ";")
			query := strings.TrimSpace(currentSQL[:semicolonIdx])

			temp := parse.Parse(query)

			switch temp.Flag {
			case "t":
				resultModel.name = temp.Query[0]
				titleDefalt = parse.ParseTableCell(temp.Query[1])
				resultModel.titles = titleDefalt.Title
			case "i":
				if resultModel.name == "" {
					return
				}
				resultModel.rows = parse.AddRowsOfModel(temp, titleDefalt)
			}

			sqlBuilder.Reset()
			if semicolonIdx+1 < len(currentSQL) {
				remaining := strings.TrimSpace(currentSQL[semicolonIdx+1:])
				if remaining != "" {
					sqlBuilder.WriteString(remaining)
				}
			}
		}
	}
	fmt.Println("Name table",resultModel.name)
	fmt.Println("Rows",resultModel.rows)
	fmt.Printf("Titles %v\n\n",resultModel.titles,)

	if err := scanner.Err(); err != nil {
		return
	}

	if resultModel.name == "" {
		return
	}
	for _, r := range resultModel.rows{
		fmt.Println(r)
	}
}



func TestAddSqlRows(t *testing.T){
	// testCell := parse.Cell{}
	readFileAddRows("./test_file/file.sql")
}
