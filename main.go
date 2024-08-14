/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"github.com/wlai-lp/bo-botflow/cmd"
	// tea "github.com/charmbracelet/bubbletea"
	"os"
	"fmt"

)

func main() {
	dir, err := os.Getwd()
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("Current working directory:", dir)

	
	cmd.Execute()
}
