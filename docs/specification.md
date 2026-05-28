
# Deta - Техническая спецификация

## 1. Общая информация
| Название проекта     | Deta Viewer                                      |
|----------------------|--------------------------------------------------|
| Тип приложения       | CLI утилита для просмотра данных                 |
| Язык реализации      | Go                                               |
| Интерфейс            | TUI (Terminal User Interface) на базе BubbleTea |
| Назначение           | Просмотр структурированных данных из файлов различных форматов |

## 2. Функциональные требования на уровне mvp
### 2.1. Поддерживаемые форматы файлов
|Формат |Расширение|
|-------|----------|
|CSV    |.csv      |
|SQL    |.sql      |
|JSON   |.json     |
|Binary|нереализованные расширения файлов|

### 2.2. Аргументы командной строки 
|Аргументы|Описание|
|---------|--------|
|-h, --help|Вывод справочной инфорпмации|
|\[filename\]|Путь к файлу для просмотра|

## 3. Архитектура проекта
```
deta/
├── cmd
│   └── main.go
├── docs
│   └── specification.md
├── examples
├── go.mod
├── go.sum
├── internal
│   ├── parseSQL
│   │   ├── insert.go
│   │   ├── parse.go
│   │   └── table.go
│   ├── readers
│   │   ├── csv_reader.go
│   │   ├── factory.go
│   │   ├── json_reader.go
│   │   └── sql_reader.go
│   │   └── binary_reader.go
│   └── parse
│       ├── binary_math.go
│       ├── binary_parse.go
│       ├── json_parse.go
├── Makefile
├── pkg
│   ├── interfaces
│   │   └── interface_reader.go
│   └── util
│       ├── interface.go
│       └── start_cmd.go
├── readme.md
├── test
│   ├── add_rows_sql_test.go
│   ├── cmd_test.go
│   ├── parse_insert_into_test.go
│   ├── parse_test.go
│   ├── read_sql_test.go
│   ├── binary_math_test.go
│   ├── binary_parse_test.go
│   └── test_file
│       ├── file.csv
│       ├── file.json
│       └── file.sql
|       └── file.py
└── utils
    ├── file_check.go
    ├── styles.go
    └── keys.go
```

## 4. Интерфейсы и структуры данных
- Интерфейс модели для пакета Bubble tea — `Reader`
  Пример реализации:
  ```golang
    type Reader interface {
        Init(path string) error // метод для инициализации
        Run() error // метод для запуска 
    }
  ```
- Структуры для хранения информации после парсинга SQL файлов
  
  Пример реализации:
  ```golang
    // Структура для хранения данных после парсинга
    type Cell struct{
        Query []string // распаршенный запрос с названием таблицы 
        Flag string // флаговая переменная для определения состояния парсинга кода
        InValue [][]string // для данных в INSERT INTO 
    }

    // Структура для хранения default значений
    type CellDefault struct{
        Title []string // столбцы таблицы
        DefaultVal map[string]string // default значения для каждого столбца
        Index string // Столбец содержащий PRIMARY KEY
    }
  ```
- Структура модели для каждого из расширений файлов — `ModelReader...`
  
  Пример реализации:
  ```golang
    type ModelReaderCSV struct {
        Path   string // путь к файлу с данными
        Title  *[]string // название таблицы
        Rows   *[][]string // данные из файла
        Table  table.Model // таблица
        Style  utils.ReaderStyles // стили для таблицы
    }
  ```
- Структуры для хранения информации после парсинга для передачи в модели расширений — `...LoadedMsg`
  
  Пример реализации:
  ```golang
    type csvLoadedMsg struct {
        title []string // хранит заголовки для колонок таблицы
        rows  [][]string // хранит данные для таблицы
    }
  ```
- Тип ошибка для каждого расширения файла — `...ErrorMsg`
  
  Пример реализации:
  ```golang
    type csvErrorMsg error
  ```

## 5. Сторонние пакеты, версии и назначения
|Пакет|	Версия|	Назначение|
|-----|-------|-----------|
|github.com/charmbracelet/bubbletea|	v1.3.10 |	TUI фреймворк|
|github.com/evertras/bubble-table|	v0.19.2 |	Табличный компонент|
|github.com/charmbracelet/lipgloss|	v1.1.0|	Стилизация|
## 6. Детали реализации
### 6.1 CSV Reader
#### Алгоритм работы
- Проверка существования файла (utils.CheckFile)
- Построчное чтение файла
- Разделение строк по разделителям: , ; | \t
- Первая строка — заголовки столбцов
- Остальные строки — данные
- Построение BubbleTea таблицы с динамической шириной столбцов
#### Особенности
- Автоматический расчет ширины столбцов
- Поддержка различных разделителей через regexp

### 6.2 SQL Reader
#### Алгоритм работы:
- Проверка существования файла
- Фильтрация комментариев (-- и #)
- Буферизация до символа ; (разделитель запросов)
- Парсинг CREATE TABLE — получение структуры
- Парсинг INSERT INTO — получение данных
- Применение DEFAULT значений и генерация PRIMARY KEY
- Построение таблицы

#### Особенности:
- Поддержка многострочных SQL запросов
- Автоматическая генерация ID для PRIMARY KEY
- Обработка DEFAULT значений для отсутствующих столбцов

### 6.3 Парсер SQL

#### ParseTableCell - извлекает из CREATE TABLE:
- Имена столбцов (без DEFAULT)
- DEFAULT значения
- PRIMARY KEY столбец
#### AddRowsOfModel - формирует итоговые строки:
- Маппинг вставляемых столбцов на позиции
- Заполнение DEFAULT значений
- Генерация PRIMARY KEY (порядковый номер строки)

### 7. Стили и оформление
Цветовая схема (палитра Kanagawa)
|Элемент|	Цвет (Lipgloss)|	Описание|
|-------|------------------|------------|
|Заголовок таблицы|	#957fb8 |	Фон заголовков|
|Строки данных| #223249 |	Фон выделенной строки|
