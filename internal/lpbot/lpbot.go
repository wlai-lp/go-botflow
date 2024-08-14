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

func GenerateMermaidChart(lpBot *LPBot) (string) {
	fmt.Println("bot id is " + lpBot.Bot.ID)
	// Generate Mermaid chart
	var mermaid strings.Builder
	mermaid.WriteString("graph TD;\n")

	// bot high level has dialogs
	dialogs := outputDialogs(lpBot)
	mermaid.WriteString(dialogs)

	subgraphs := outputSubgraphs(lpBot)
	fmt.Printf("contains elements %v\n", len(subgraphs))
	for _, sub := range subgraphs {
		mermaid.WriteString(sub)
	}
	// Output the Mermaid chart
	fmt.Println("========== mermaid =========")
	// fmt.FormatString()
	result := fmt.Sprintln(mermaid.String())
	return result
}


func outputSubgraphs(lpBot *LPBot) []string {

	// Create a slice with the same length as lpBot.Groups
	strSlice := make([]string, len(lpBot.Groups))

	g := lpBot.Groups
	cm := lpBot.ConversationMessage

	// var cmessages []ConversationMessage
	// // Group by "group"
	// groupedMessages := make(map[string][]ConversationMessage)
	// for _, message := range cmessages {
	// 	groupedMessages[message.Group] = append(groupedMessages[message.Group], message)
	// }

	cmMap := make(map[string][]ConversationMessage)

	// create a map[dialog id] []ConversationMessages
	for _, el := range cm {
		cmMap[el.Group] = append(cmMap[el.Group], el)

	}

	// for each dialog, build subgraph
	for index, element := range g {
		var subgraph strings.Builder
		subgraph.WriteString(fmt.Sprintf("   subgraph %v[%v]\n", element.ID, element.Name))
		// TODO: build subgraph flow
		// input is a map[groupid][]convomessage
		msgArray := cmMap[element.ID]
		for _, msg := range msgArray {
			subgraph.WriteString(fmt.Sprintf("      %v[%v]\n", msg.ID, msg.Name))
			if msg.NextMessageId != "" {
				fmt.Println("next message id is not nil " + msg.NextMessageId)
				subgraph.WriteString(fmt.Sprintf("      %v --> %v\n", msg.ID, msg.NextMessageId))
			}
		}

		subgraph.WriteString("   end\n\n")
		strSlice[index] = subgraph.String()
	}

	// output := [3]string{"abc", "efg", "abc"}
	// count := len(lpBot.Groups)
	// fmt.Println(count)
	// return output[:]
	return strSlice[:]

}

func outputDialogs(lpBot *LPBot) string {
	var dialogs strings.Builder
	dialogs.WriteString("   " + lpBot.Bot.ID + "[" + lpBot.Bot.Name + "]\n")
	g := lpBot.Groups
	for index, element := range g {
		fmt.Printf("At index %d, value is %s\n", index, element.DialogType)
		dialogs.WriteString(fmt.Sprintf("   %v -->|has| %v[%v]\n", lpBot.Bot.ID, element.ID, element.Name))
	}
	dialogs.WriteString("\n")
	return dialogs.String()
}

/**
bot json structure important

top level is bot

bot.conversationMessages[] is array
- each conversationMessage is an interaction

each conversationMessage has
- id (unique to that interaction, use this to id mermaid node)
- name (use this to label mermaid node)
- status "active"  // to be implemented
- group (belongs to the dialog)
- content
	contentType ...// to be implemented
- nextMessageId // to create the link relationship in mermaid

**/