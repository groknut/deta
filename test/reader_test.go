package test

import (
	"context"
	"deta/readers"
	"path/filepath"
	"testing"
)

func TestReaderCSV(t *testing.T) {
	path, _ := filepath.Abs(".")
	path = filepath.Join(path,"file.csv")
	readContext, readCancel := context.WithCancel(context.Background())

	ctx := ro

	title := make([]string,0)
	rows :=  make([][]string,0)
	model := readers.ModelReaderCSV{Path: path, 
							Ctx: ctx,
							Title: &title,
							Rows: &rows}
}