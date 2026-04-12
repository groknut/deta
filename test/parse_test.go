package test

import(
	"testing"
	"fmt"
	parse "deta/internal/parseSQL"
)

func TestParse(t *testing.T){
	fmt.Println(parse.Parse("CREATE TABLE orders (id PRIMARY KEY, user_id INTEGER, total_amount DECIMAL(10,2) );"))
	fmt.Println(parse.Parse("CREATE TABLE user (id PRIMARY KEY, user_id INTEGER, total_amount INTEGER);"))

}