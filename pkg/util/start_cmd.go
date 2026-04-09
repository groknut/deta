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

	filename := os.Args[1]

	if len(os.Args) < 2 || slices.Contains(HELP_ARGS, filename) {
		fmt.Print(HELP_MESSAGE)
		return nil
	}

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
