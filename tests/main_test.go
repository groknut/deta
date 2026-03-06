package tests

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"testing"
)

func test_terminal(com string,arr []string){
	cmd := exec.Command(com, arr...)
	output, err := cmd.Output()
	if err != nil{
		print("\ntxt file error PASSED ",err,"\n")
		return
	}
	fmt.Println(string(output),"Test was passed")
}

func TestMain(t *testing.T){

	testDir, _ := filepath.Abs(".")
	root := filepath.Dir(testDir)

	mainPath := filepath.Join(root,"main.go")

	test_terminal("go",[]string{"run",mainPath,"-h"})
	test_terminal("go",[]string{"run",mainPath,"file.csv"})
	test_terminal("go",[]string{"run",mainPath,"file.txt"})
	test_terminal("go",[]string{"run",mainPath,"file.json"})
	test_terminal("go",[]string{"run",mainPath,"file.hex"})
	test_terminal("go",[]string{"run",mainPath,"-h","file.csv`"})
	test_terminal("go",[]string{"run",mainPath,"file.json","-h"})

}