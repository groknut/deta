package tests

import (
	"fmt"
	"os/exec"
	"testing"
)

func test_terminal(com string,arr []string){
	cmd := exec.Command(com, arr...)
	output, err := cmd.Output()
	if err != nil{
		print("Error",err)
		return
	}
	fmt.Println(string(output),"Test was passed")
}

func TestMain(t *testing.T){
	test_terminal("go",[]string{"run",`C:\Users\User\Desktop\рпо\дз_второй_курс\fossdev-project\main.go`,"-h"})
	test_terminal("go",[]string{"run",`C:\Users\User\Desktop\рпо\дз_второй_курс\fossdev-project\main.go`,"file.csv"})
	test_terminal("go",[]string{"run",`C:\Users\User\Desktop\рпо\дз_второй_курс\fossdev-project\main.go`,"file.txt"})
	test_terminal("go",[]string{"run",`C:\Users\User\Desktop\рпо\дз_второй_курс\fossdev-project\main.go`,"file.json"})
	test_terminal("go",[]string{"run",`C:\Users\User\Desktop\рпо\дз_второй_курс\fossdev-project\main.go`,"file.hex"})
	test_terminal("go",[]string{"run",`C:\Users\User\Desktop\рпо\дз_второй_курс\fossdev-project\main.go`,"-h","file.csv`"})
	test_terminal("go",[]string{"run",`C:\Users\User\Desktop\рпо\дз_второй_курс\fossdev-project\main.go`,"file.json","-h"})

}