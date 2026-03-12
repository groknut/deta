package util

import (
	"deta/readers"
	"errors"
	"fmt"
	"os"
	"strings"
	"path/filepath"
)


//Функция для запуска утилиты
func StartUtil() error{
	quantityArgs := len(os.Args)
	// fmt.Println(quantityArgs)
	if quantityArgs < 2{
		return errors.New("File not transferred")
	}
	if os.Args[1] == "-h" || os.Args[1] =="--help"{
		fmt.Println(`Commands...`)
		return nil
	}

	filename := os.Args[1]
    ext := filepath.Ext(filename)
    typeFile := strings.TrimPrefix(ext, ".")

	//Вызывать методы для обработки файлов вызывать здесь
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