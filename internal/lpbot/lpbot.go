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

func GenerateMermaidChart(lpBot *LPBot) {
	fmt.Println("bot id is " + lpBot.Bot.ID)
	// Generate Mermaid chart
	var mermaid strings.Builder
	mermaid.WriteString("graph TD;\n")

	// bot high level has dialogs
	dialogs := outputDialogs(lpBot)
	mermaid.WriteString(dialogs)

	subgraphs := outputSubgraphs(lpBot)
	fmt.Println(len(subgraphs))
	// Output the Mermaid chart
	fmt.Println(mermaid.String())
}

func outputSubgraphs(lpBot *LPBot) []string {

	// Create a slice with the same length as lpBot.Groups
	strSlice := make([]string, len(lpBot.Groups))

	g := lpBot.Groups
	// cm := lpBot.ConversationMessage

	// cmMap := make(map[string][]ConversationMessage)

	// for _, el := range cm {
	// 	cmMap[el.Group] =

	// }

	for index, element := range g {
		var subgraph strings.Builder
		subgraph.WriteString(fmt.Sprintf("   subgraph %v[%v]", element.ID, element.Name))
		strSlice[index] = subgraph.String()
		subgraph.WriteString("   end")
	}

	output := [3]string{"abc", "efg", "abc"}
	count := len(lpBot.Groups)
	fmt.Println(count)
	return output[:]

}

func outputDialogs(lpBot *LPBot) string {
	var dialogs strings.Builder
	dialogs.WriteString("   " + lpBot.Bot.ID + "[" + lpBot.Bot.Name + "]\n")
	g := lpBot.Groups
	for index, element := range g {
		fmt.Printf("At index %d, value is %s\n", index, element.DialogType)
		dialogs.WriteString(fmt.Sprintf("   %v -->|has| %v[%v]\n", lpBot.Bot.ID, element.ID, element.Name))
	}

	return dialogs.String()
}
