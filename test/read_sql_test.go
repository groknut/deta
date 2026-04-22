package test

import(
	"bufio"
	"testing"
	"fmt"
	"strings"
	"os"
)

func read(path string) string{
	file, err := os.Open(path)
	if err != nil {
		return "File doesn't exists"
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if err != nil {
		return "Error of reading"
	}
	var sqlQuery string
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if strings.Contains(line, ";"){
			sqlQuery += line
			
		} else{
			sqlQuery += line + " "
		}

	}
	return sqlQuery
}

func TestSQLRead(t *testing.T){
	fmt.Println(read("./test_file/file.sql"))
}