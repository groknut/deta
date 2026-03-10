package util

import (
	"deta/readers"
	"errors"
	"fmt"
	"os"
	"slices"
	"strings"
	// "path/filepath"
)

// Функция для запуска утилиты
func StartUtil() error {

	if len(os.Args) < 2 || slices.Contains(HELP_ARGS, os.Args[1]) {
		fmt.Print(HELP_MESSAGE)
		return nil
	}

	var typeFile string
	splitFileName := strings.Split(os.Args[1], ".")

	lenghtNameFile := len(splitFileName)
	if lenghtNameFile == 2 {
		typeFile = splitFileName[1]
	} else {
		typeFile = splitFileName[lenghtNameFile-1]
	}

	//Вызывать методы для обработки файлов вызывать здесь
	switch typeFile {
	case "csv":
		readers.StartReaderCSV(os.Args[1])
	case "json":
		fmt.Println(typeFile)
	case "hex":
		fmt.Println(typeFile)

	default:
		return errors.New("Undefine file extension")
	}

	return nil
}
