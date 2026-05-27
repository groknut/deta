package util

import (
	"deta/internal/readers"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
	"path/filepath"
)

func StartUtil() error{

	if len(os.Args) < 2 || slices.Contains(HELP_ARGS, os.Args[1]) {
		fmt.Print(HELP_MESSAGE)
		return nil
	}

	filename := os.Args[1]
    ext := filepath.Ext(filename)
    typeFile := strings.TrimPrefix(ext, ".")

	reader, err := readers.GetReader(typeFile)
	if err != nil{
		// Вызов дефолтного пейджера
		return nil
	}
	if err := reader.Init(os.Args[1]);err != nil{
		return errors.New("Undefine error")
	}

	if err := reader.Run();err != nil{
		return errors.New("Failed to start reader")
	}

	return nil
}
