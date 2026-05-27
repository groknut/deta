package readers

import "deta/pkg/interfaces"

// Словарь для вызова структур
var FactoryTypeFile = map[string] func() interfaces.Reader{
	"csv": func() interfaces.Reader { return NewCSVReader() },
	"sql": func() interfaces.Reader { return NewSQLReader() },
}

// Подает метод для вызова чтения файлов
func GetReader(typeFile string) (interfaces.Reader, error){
	factory, exists := FactoryTypeFile[typeFile]
	if !exists {
		return NewBinaryReader(), nil
	}

	return factory(), nil
}
