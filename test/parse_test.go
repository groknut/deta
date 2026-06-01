package test

import(
	"testing"
	"fmt"
	parse "github.com/groknut/deta/internal/parseSQL"
)

func testCreateTable(str string) bool {
	temp := parse.Parse(str)
	return !(len(temp.Query) == 0)
}

func TestParse(t *testing.T){
	i := 0
	fmt.Println()
	if testCreateTable("INSERT INTO users (name, email) VALUES ('Alice Johnson', 'alice@com'), ('Bob Mil', 'Bob@Com');"){
		i++
	}
	if testCreateTable("INSERT INTO users (name,  email) VALUES ('Alice Johnson', 'alice@com'), ('Bob Mil', 'Bob@Com');"){
		i++
	}

	if testCreateTable("INSERT INTO users (name, email) VALUES ('Alice Johnson', 'alice@com'), ('Bob Mil', 'Bob@Com');"){
		i++
	}
	if i == 3{
		fmt.Println("TestCreateTable PASSED")
		return
	}
	fmt.Println("TestCreateTable didn't PASSED")


}
