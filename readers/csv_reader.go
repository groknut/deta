package readers

//Модель для BubbleTea для csv файлов
type readerCSV struct{
	cursor int
	rows chan []string
	titles chan string	
}


//начало работы csv файла
func StartCSV() {

}

//чтение csv файла в отдельной горутине и передача данных в модель для отображения
func(r *readerCSV) ReadCSV() {

}