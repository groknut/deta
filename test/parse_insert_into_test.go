package test

import(
	"testing"
	"fmt"
	parse "deta/internal/parseSQL"
)

func TestInsertInto(t *testing.T){
	fmt.Println(parse.Parse("INSERT INTO users (name, email) VALUES ('Alice Johnson', 'alice@com');"))
	fmt.Println(parse.Parse("INSERT INTO users (name,  email) VALUES ('Alice Johnson', 'alice@com');").Query[0])
	fmt.Println(parse.Parse("INSERT INTO users (name,   email) VALUES ('Alice Johnson', 'alice@com');").Query[1])
	fmt.Println(parse.Parse("INSERT INTO users (name, email) VALUES ('Alice Johnson', 'alice@com');").InValue)
}