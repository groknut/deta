package util

import (
	"errors"
	"fmt"
	"os"
	"strings"
	// "path/filepath"
	// "util/readers"
)

//Функция для запуска утилиты
func StartUtil() error{
	quantityArgs := len(os.Args)
	fmt.Println(quantityArgs)
	if quantityArgs < 2{
		return errors.New("File not transferred")
	}

	var typeFile string
	splitFileName := strings.Split(os.Args[1],".")
	lenghtNameFile := len(splitFileName) 
	if lenghtNameFile == 2{
		typeFile = splitFileName[1]
	}else{
		typeFile = splitFileName[lenghtNameFile-1]
	}

	
	switch typeFile{
	case "csv":
	case "json":

	default:
		return errors.New("Undefine file extension")
	}
	
	return nil
}