package main

import (
	deta "github.com/groknut/deta/pkg/util"
	"fmt"
	"os"
)

func main() {
	if err := deta.StartUtil(); err != nil{
		fmt.Println(err)
		os.Exit(1)
	}
	// fmt.Println("Hello, world!")
}
