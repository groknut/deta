package interfaces



// Интерфейс читатель
type Reader interface{
	Init(path string)	error
	Run() error
}

type ReaderFactory func() Reader
