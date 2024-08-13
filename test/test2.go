package main

import (
	"encoding/json"
	"fmt"
)

// Message represents the structure of each JSON object in the array.
type Message struct {
	ID                string      `json:"id"`
	ChatBotId         string      `json:"chatBotId"`
	UserInputRequired bool        `json:"userInputRequired"`
	Name              string      `json:"name"`
	Type              string      `json:"type"`
	Content           Content     `json:"content"`
	Group             string      `json:"group"`
	Status            string      `json:"status"`
	PrevMessageId     string      `json:"prevMessageId,omitempty"`
	ResponseMatches   []Response  `json:"responseMatches"`
	InteractionType   string      `json:"interactionType"`
	Pattern           []string    `json:"pattern,omitempty"`
	PreProcessMessage string      `json:"preProcessMessage,omitempty"`
	NextMessageId     string      `json:"nextMessageId,omitempty"`
}

type Content struct {
	ContentType string  `json:"contentType"`
	Results     Results `json:"results"`
}

type Results struct {
	Type string `json:"type"`
	Tile Tile   `json:"tile"`
}

type Tile struct {
	TileData []TileData `json:"tileData"`
}

type TileData struct {
	Text          string   `json:"text"`
	Buttons       []string `json:"buttons"`
	QuickReplyList []string `json:"quickReplyList"`
}

type Response struct {
	Conditions            []string `json:"conditions"`
	ContextConditions     []string `json:"contextConditions"`
	Action                Action   `json:"action"`
	ContextDataVariables  []string `json:"contextDataVariables"`
}

type Action struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

func main() {
	// JSON input
	jsonData := `
	[
		{
			"id": "5c7fdcf600d9ddbcbc56b60031bf65bec9de253b",
			"chatBotId": "b3f30d7f-22b1-4e53-9f43-c3d2f927a3bd",
			"userInputRequired": false,
			"name": "fallback_starter",
			"type": "BOT_MESSAGE",
			"content": {
				"contentType": "STATIC",
				"results": {
					"type": "TEXT",
					"tile": {
						"tileData": [
							{
								"text": "Sorry, I am not able to understand.",
								"buttons": [],
								"quickReplyList": []
							}
						]
					}
				}
			},
			"group": "b31a6e90-8dfc-4872-862a-35ac85e88966",
			"status": "ACTIVE",
			"responseMatches": [
				{
					"conditions": [],
					"contextConditions": [],
					"action": {
						"name": "INTERACTION",
						"value": "next"
					},
					"contextDataVariables": []
				}
			],
			"interactionType": "TEXT"
		},
		{
			"id": "7450135e679558cee81370e271dc44c5b89dd19b",
			"chatBotId": "b3f30d7f-22b1-4e53-9f43-c3d2f927a3bd",
			"userInputRequired": false,
			"name": "text_message",
			"type": "BOT_MESSAGE",
			"content": {
				"contentType": "STATIC",
				"results": {
					"type": "TEXT",
					"tile": {
						"tileData": [
							{
								"text": "Hi there! Thanks for coming!",
								"buttons": [],
								"quickReplyList": []
							}
						]
					}
				}
			},
			"group": "c8e5d31e-42f4-4155-989e-d3f16ee36f15",
			"status": "ACTIVE",
			"prevMessageId": "7942ccd028d5d1a2de30c5269dc88ee8993a771d",
			"responseMatches": [
				{
					"conditions": [],
					"contextConditions": [],
					"action": {
						"name": "INTERACTION",
						"value": "next"
					},
					"contextDataVariables": []
				}
			],
			"interactionType": "TEXT"
		},
		{
			"id": "7942ccd028d5d1a2de30c5269dc88ee8993a771d",
			"chatBotId": "b3f30d7f-22b1-4e53-9f43-c3d2f927a3bd",
			"userInputRequired": false,
			"name": "dialog_starter",
			"pattern": [
				"hi",
				"hello",
				"howdy",
				"(hi|hello|hey)*"
			],
			"type": "BOT_MESSAGE",
			"content": {
				"contentType": "STATIC",
				"results": {
					"type": "TEXT",
					"tile": {
						"tileData": [
							{
								"text": "Hello!!",
								"buttons": [],
								"quickReplyList": []
							}
						]
					}
				}
			},
			"preProcessMessage": "var d = new Date(); \nif (botContext.getBotVariable('lastVisited') != null) { \n    botContext.setTriggerNextMessage('Welcome Back');\n} else { \n    botContext.setBotVariable('lastVisited',d,false,true); \n}",
			"group": "c8e5d31e-42f4-4155-989e-d3f16ee36f15",
			"status": "ACTIVE",
			"nextMessageId": "7450135e679558cee81370e271dc44c5b89dd19b",
			"responseMatches": [
				{
					"conditions": [],
					"contextConditions": [],
					"action": {
						"name": "INTERACTION",
						"value": "next"
					},
					"contextDataVariables": []
				}
			],
			"interactionType": "DIALOG_STARTER"
		}
	]`

	// Parse JSON data
	var messages []Message
	err := json.Unmarshal([]byte(jsonData), &messages)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Group by "group"
	groupedMessages := make(map[string][]Message)
	for _, message := range messages {
		groupedMessages[message.Group] = append(groupedMessages[message.Group], message)
	}

	// Print grouped messages
	for group, msgs := range groupedMessages {
		fmt.Printf("Group: %s\n", group)
		for _, msg := range msgs {
			fmt.Printf("  ID: %s, Name: %s, Type: %s\n", msg.ID, msg.Name, msg.Type)
		}
	}
}

