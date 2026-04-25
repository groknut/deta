package test

import(
	"testing"
	"fmt"
	parse "deta/internal/parseSQL"
)

func TestInsertInto(t *testing.T){
	fmt.Println(parse.Parse("INSERT INTO users (name, email) VALUES ('Alice Johnson', 'alice@com'), ('Bob Mil', 'Bob@Com');"))
	fmt.Println(parse.Parse("INSERT INTO users (name,  email) VALUES ('Alice Johnson', 'alice@com'), ('Bob Mil', 'Bob@Com');;").Query[0])
	fmt.Println(parse.Parse("INSERT INTO users (name,   email) VALUES ('Alice Johnson', 'alice@com'), ('Bob Mil', 'Bob@Com');;").Query[1])
	fmt.Println(parse.Parse("INSERT INTO users (name, email) VALUES ('Alice Johnson', 'alice@com'), ('Bob Mil', 'Bob@Com');;").InValue[1])
}