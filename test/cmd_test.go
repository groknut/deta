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
		print("Test didn't PASSED",err,"\n")
		return
	}
	fmt.Println(string(output),"Test was PASSED")
}

func TestMain(t *testing.T){

	testDir, _ := filepath.Abs(".")
	root := filepath.Dir(testDir)
	mainPath := filepath.Join(root,"main.go")
	test_terminal("go",[]string{"run",mainPath,"-h"})
	test_terminal("go",[]string{"run",mainPath,"test/test_file/file.csv"})
	test_terminal("go",[]string{"run",mainPath,"test/test_file/file.sql"})
	test_terminal("go",[]string{"run",mainPath,"-h","test/test_test/file.csv`"})
	
}