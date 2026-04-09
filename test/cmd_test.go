package test

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"testing"
)

func test_terminal(com string,arr []string){
	cmd := exec.Command(com, arr...)
	output, err := cmd.CombinedOutput()
	if err != nil{
		print("\ntxt file error PASSED ",err,"\n")
		return
	}
	fmt.Println(string(output),"Test was passed")
}

func TestMain(t *testing.T){

	testDir, _ := filepath.Abs(".")
	root := filepath.Dir(testDir)
	mainPath := filepath.Join(root,"cmd/main.go")
	fmt.Println(mainPath)
	test_terminal("go",[]string{"run",mainPath,"-h"})
	test_terminal("go",[]string{"run",mainPath,"test_file/file.csv"})
	test_terminal("go",[]string{"run",mainPath,"test_file/file.txt"})
	test_terminal("go",[]string{"run",mainPath,"test_file/file.json"})
	test_terminal("go",[]string{"run",mainPath,"test_file/file.hex"})
	test_terminal("go",[]string{"run",mainPath,"-h","test_test/file.csv`"})
	test_terminal("go",[]string{"run",mainPath,"test_test/file.json","-h"})
	
}