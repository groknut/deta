package readers

import (
	"deta/pkg/interfaces"
	"errors"
	// "deta/readers"
)

// Словарь для вызова структур
var FactoryTypeFile = map[string] func() interfaces.Reader{
	"csv": func() interfaces.Reader { return NewCSVReader() },
}

// Подает метод для вызова чтения файлов
func GetReader(typeFile string) (interfaces.Reader, error){
	factory, exists := FactoryTypeFile[typeFile]
	if !exists{
		return nil, errors.New("Error type file")
	}

	return factory(), nil
}