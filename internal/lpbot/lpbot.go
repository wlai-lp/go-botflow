package lpbot

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func LoadBotFile(f string) (*LPBot, error) {
	fmt.Printf("load bot file %v\n", f)
		
	data, err := os.ReadFile(f)
	if err != nil {
		return nil, err
	}
	var bot LPBot
	err = json.Unmarshal(data, &bot)
	if err != nil {
		return nil, err
	}
	fmt.Println("hello from lp bot")
	return &bot, nil
}

func GenerateMermaidChart(lpBot *LPBot){
	fmt.Println("bot id is " + lpBot.Bot.ID)
	// Generate Mermaid chart
	var mermaid strings.Builder
	mermaid.WriteString("graph TD;\n")

	// bot high level has dialogs
	dialogs := outputDialogs(lpBot)
	mermaid.WriteString(dialogs)
	// Output the Mermaid chart
	fmt.Println(mermaid.String())
}

func outputDialogs(lpBot *LPBot) (string){
	var dialogs strings.Builder
	dialogs.WriteString("   " + lpBot.Bot.ID + "[" + lpBot.Bot.Name + "]\n")
	g := lpBot.Groups
	for index, element := range g {
		fmt.Printf("At index %d, value is %s\n", index, element.DialogType)
		dialogs.WriteString(fmt.Sprintf("   %v -->|has| %v[%v]\n", lpBot.Bot.ID, element.ID, element.Name))
	}
	
	return dialogs.String()
}