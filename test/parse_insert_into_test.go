package test

import(
	"testing"
	"fmt"
	parse "github.com/groknut/deta/internal/parseSQL"
)

func testParse(str string) bool{
	temp := parse.Parse(str)
	return !(len(temp.Query) == 0)

}

func TestInsertInto(t *testing.T){
	i := 0
	fmt.Println()
	if testParse("INSERT INTO users (name, email) VALUES ('Alice Johnson', 'alice@com'), ('Bob Mil', 'Bob@Com');"){
		i++
	}
	if testParse("INSERT INTO users (name,  email) VALUES ('Alice Johnson', 'alice@com'), ('Bob Mil', 'Bob@Com');"){
		i++
	}

	if testParse("INSERT INTO users (name, email) VALUES ('Alice Johnson', 'alice@com'), ('Bob Mil', 'Bob@Com');"){
		i++
	}
	if i == 3{
		fmt.Println("TestInsertInto PASSED")
		return
	}
	fmt.Println("TestInsertInto didn't PASSED")

}
