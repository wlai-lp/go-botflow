package lpbot

import (
	"fmt"
	"os"
	"testing"
)

func TestLPBotLoad(t *testing.T){
	dir, _ := os.Getwd()
	botfile := "BotFlow.json"
	fileFullPath := dir + string(os.PathSeparator) + "testdata" + string(os.PathSeparator) + botfile
	
	_, err := LoadBotFile(fileFullPath)
	if err != nil {
		t.Fatal("LP Bot Load returned error")
	}
	fmt.Println("testing done")
}

func TestLPBotLoadAndGenerate(t *testing.T){
	dir, _ := os.Getwd()
	botfile := "BotFlow.json"
	fileFullPath := dir + string(os.PathSeparator) + "testdata" + string(os.PathSeparator) + botfile
	
	bot, err := LoadBotFile(fileFullPath)
	if err != nil {
		t.Fatal("LP Bot Load returned error")
	}
	result := GenerateMermaidChart(bot)
	fmt.Println("testing done \n" + result)
}
