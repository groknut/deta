package test

import (
    "fmt"
    "os"
    "os/exec"
    "path/filepath"
    "testing"
    "time"
)

func testModelTea(com string, arr []string) {
    cmd := exec.Command(com, arr...)
    
    cmd.Stdin = os.Stdin
    cmd.Stdout = os.Stdout
    cmd.Stderr = os.Stderr
    
    err := cmd.Start()
    if err != nil {
        fmt.Printf("Ошибка запуска: %v\n", err)
        return
    }
    
    time.Sleep(2 * time.Second)
    err = cmd.Wait()
    if err != nil {
        fmt.Printf("Программа завершилась с ошибкой: %v\n", err)
    }
}

func TestMain(t *testing.T) {
    if testing.Short() {
        t.Skip("Skipping TUI test in short mode")
    }
    
    testDir, _ := filepath.Abs(".")
    root := filepath.Dir(testDir)
    mainPath := filepath.Join(root, "main.go")
    csvPath := filepath.Join(root, "file.csv")
    fmt.Printf("Запуск TUI с файлом: %s\n", csvPath)
    
    testModelTea("go", []string{"run", mainPath, csvPath})
}