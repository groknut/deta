# Настройки вывода

Output assets/demo.gif
Set WindowBar Colorful
Set FontSize 20
Set Width 1200
Set Height 600

Hide
Type "prompt $g "
Enter
Type "cls"
Enter

Show
Type "deta.exe -h"
Enter

# Смотрим на -help
Sleep 1s

Type "deta.exe test/test_file/file.csv"
Enter
Sleep 1s
Down
Sleep 1s
Up

Type "q"

Type "deta.exe test/test_file/file.json"
Enter
Sleep 1s
Down
Sleep 1s
Down
Sleep 1s
Right
Sleep 1s
Left
Sleep 1s
Down
Sleep 1s
Right
Sleep 1s
Left
Sleep 1s
Down
Sleep 500ms
Down

Type "q"

Type "deta.exe test/test_file/file.sql"
Enter
Sleep 1s
Down
Sleep 1s
Up

Type "q"

Type "deta.exe readme.md"
Enter
Sleep 1s
Down
Sleep 500ms
Down
Sleep 500ms
Down
Sleep 500ms
Down
Sleep 500ms
Down


# Выходим из программы (замените 'q' на вашу клавишу выхода, если она другая)
Type "q"
Sleep 500ms
