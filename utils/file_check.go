package utils

import "os"

func CheckFile(path string) error{
    _, err := os.Stat(path)
    if err != nil{
        if os.IsNotExist(err){
            return err
        }
    }
    return nil
}
