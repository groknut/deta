package test

import (
	"context"
	"deta/readers"
	"fmt"
	"path/filepath"
	"testing"
)

func testReaderCSV() {
	path, _ := filepath.Abs(".")
	path = filepath.Join(path,"file.csv")
	readContext, readCancel := context.WithCancel(context.Background())

	

	title := make([]string,0)
	rows :=  make([][]string,0)
	model := readers.ModelReaderCSV{Path: path, 
							Ctx: readContext,
							Title: &title,
							Rows: &rows,
							CtxCancel: readCancel}
	model.ReadCSV()
	fmt.Println(model.Title)
	fmt.Println(model.Rows)
}

func testModelTea(){
	path, _ := filepath.Abs(".")
	path = filepath.Join(path,"file.csv")
	readContext, readCancel := context.WithCancel(context.Background())

	readers.StartReaderCSV(readContext,readCancel,path)
}

func TestMain(t *testing.T){
	// testReaderCSV()
	testModelTea()
}
